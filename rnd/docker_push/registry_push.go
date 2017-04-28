package main

// TODO: commit tag, event type and branch flags and logic
// TODO: create one each subscripts for gcr and docker

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Workflow
	Registry
}

var configFile string

func init() {
	const (
		defaultConfigFile = "cicd.yaml"
		configFileUsage   = "configuration file containing project workflow values"
	)
	flag.StringVar(&configFile, "config", defaultConfigFile, configFileUsage)
	flag.StringVar(&configFile, "c", defaultConfigFile, configFileUsage)
}

func tag(r Registrator, tag string) (string, error) {
	return r.Tag(tag)
}

func push(r Registrator) (string, error) {
	return r.Push()
}

func isValid(r Registrator) bool {
	return r.IsValid()
}

func main() {

	flag.Parse()

	// read in project config file
	yamlInput, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// parse yaml into Config object
	cfg := Config{}
	err = yaml.Unmarshal([]byte(yamlInput), &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// point to active registry (docker, gcr, ...)
	var activeRegistry interface{}

	switch cfg.Workflow.Registry {
	case "gcr":
		activeRegistry = &cfg.Registry.GCRRegistry
	case "docker":
		activeRegistry = &cfg.Registry.DockerRegistry
	default:
		fmt.Fprintf(os.Stderr, "error: unsupported registry: %v\n", cfg.Workflow.Registry)
		os.Exit(1)
	}

	// assert type Registrator
	ar := activeRegistry.(Registrator)

	// validate registry has required values
	if ok := isValid(ar); !ok {
		fmt.Fprintf(os.Stderr, "error: missing registry url from configuration: %#v\n", ar)
		os.Exit(1)
	}

	// tag images
	var result string
	result, err = tag(ar, "master")
	fmt.Println("Tag Result:", result)

	// push images
	result, err = push(ar.(Registrator))
	fmt.Println("Push Result:", result)

}

func debugYAML(yamlInput []byte, cfg Config) {

	fmt.Println("DEBUG yamlInput:", yamlInput)

	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(yamlInput), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err := yaml.Marshal(&cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
