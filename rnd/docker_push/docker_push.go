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
	App
	Workflow
	Github
	Registry map[string]map[string]Platform
	Fruits   map[string][]string
}

// echo "DOCKER_REPO: $DOCKER_REPO"
// echo "REPO_TARGET: $REPO_TARGET"
// echo "DOCKER_COMMIT_TAG: $DOCKER_COMMIT_TAG"
// echo "BRANCH_REGEX: $BRANCH_REGEX"
// echo "TRAVIS_BRANCH: $TRAVIS_BRANCH"
// echo "TRAVIS_EVENT_TYPE: $TRAVIS_EVENT_TYPE"

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

	// fmt.Printf("Config ==> %#v\n", cfg)

	fmt.Printf("Config.App ==> %#v\n", cfg.App)
	fmt.Printf("Config.Workflow ==> %#v\n", cfg.Workflow)
	fmt.Printf("Config.Github ==> %#v\n", cfg.Github)
	fmt.Printf("Config.Registry ==> %#v\n", cfg.Registry)
	// fmt.Printf("Config.Fruits ==> %#v\n", cfg.Fruits)
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
