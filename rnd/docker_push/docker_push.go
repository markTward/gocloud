package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Registry
}

const (
	defaultConfigFile = "cicd.yaml"
)

func main() {

	// setup flags
	configFile := flag.String("config", defaultConfigFile, "configuration file containing project workflow values")
	flag.Parse()

	yamlInput, err := ioutil.ReadFile(*configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	cfg := Config{}

	err = yaml.Unmarshal([]byte(yamlInput), &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("Config.Registry ==> %v\n", cfg.Registry)

	switch cfg.Registry.Name {
	case "gcr":
		fmt.Println("pushing to GCR!!!", cfg.Registry)
	case "docker":
		fmt.Println("pushing to DOCKER!!!", cfg.Registry)
	default:
		fmt.Println("unknown Registry :-()")
	}

	// debugYAML(yamlInput, cfg)
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
