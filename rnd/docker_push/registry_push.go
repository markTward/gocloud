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

const (
	defaultConfigFile = "cicd.yaml"
	configFileUsage   = "configuration file containing project workflow values"
)

var configFile string

func init() {

	flag.StringVar(&configFile, "config", defaultConfigFile, configFileUsage)
	flag.StringVar(&configFile, "c", defaultConfigFile, configFileUsage)

}

func tag(r Registrator, tag string) (string, error) {
	return r.Tag(tag)
}

func push(r Registrator) (string, error) {
	return r.Push()
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

	// tag and push images
	var result string
	result, err = tag(activeRegistry.(Registrator), "master")
	fmt.Println("Tag Result:", result)

	result, err = push(activeRegistry.(Registrator))
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
