package main

// TODO: commit tag, event type and branch flags and logic
// TODO: create one each subscripts for gcr and docker

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Workflow
	Registry
}

var configFile, buildTag, eventType, branch *string
var pr *int

func init() {
	const (
		defaultConfigFile = "cicd.yaml"
		configFileUsage   = "configuration file containing project workflow values (default: cicd.yaml)"
		buildTagUsage     = "existing image tag used as basis for further tags (required)"
		eventTypeUsage    = "build event type from list: push, pull_request (default push)"
		branchTypeUsage   = "build branch (required)"
		prUsage           = "pull request number (required when event type is pr)"
	)
	configFile = flag.String("config", defaultConfigFile, configFileUsage)
	buildTag = flag.String("tag", "", buildTagUsage)
	eventType = flag.String("eventType", "push", eventTypeUsage)
	branch = flag.String("branch", "", branchTypeUsage)
	pr = flag.Int("pr", 0, prUsage)
}

func tag(url string, tag string, event string, branch string, pr int) ([]string, error) {
	var images []string
	// build tag image
	image := url + ":" + tag
	images = append(images, image)

	switch event {
	case "push":
		images = append(images, url+":"+branch)
		if branch == "master" {
			images = append(images, url+":latest")
		}
	case "pull_request":
		images = append(images, url+":PR-"+strconv.Itoa(pr))
	}

	return images, nil
}

func push(r Registrator) (string, error) {
	return r.Push()
}

func isRegistryValid(r Registrator) bool {
	return r.IsRegistryValid()
}

func main() {

	flag.Parse()

	// read in project config file
	yamlInput, err := ioutil.ReadFile(*configFile)
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

	// validate cli flags
	if *buildTag == "" {
		fmt.Fprintf(os.Stderr, "error: build tag a required value; use --tag option\n")
		os.Exit(1)
	}

	if *branch == "" {
		fmt.Fprintf(os.Stderr, "error: build branch a required value; use --branch option\n")
		os.Exit(1)
	}

	if *eventType == "pull_request" && *pr == 0 {
		fmt.Fprintf(os.Stderr, "error: event type pull_request requires a PR number; use --pr option\n")
		os.Exit(1)
	}

	// point to active registry (docker, gcr, ...)
	var activeRegistry interface{}
	var url string

	switch cfg.Workflow.Registry {
	case "gcr":
		activeRegistry = &cfg.Registry.GCRRegistry
		url = cfg.Registry.GCRRegistry.Url
	case "docker":
		activeRegistry = &cfg.Registry.DockerRegistry
		url = cfg.Registry.DockerRegistry.Url
	default:
		fmt.Fprintf(os.Stderr, "error: unsupported registry: %v\n", cfg.Workflow.Registry)
		os.Exit(1)
	}

	// assert type Registrator
	ar := activeRegistry.(Registrator)

	// validate registry has required values
	if ok := isRegistryValid(ar); !ok {
		fmt.Fprintf(os.Stderr, "error: missing registry url from configuration: %#v\n", ar)
		os.Exit(1)
	}

	var images []string
	images, _ = tag(url, *buildTag, *eventType, *branch, *pr)
	fmt.Println("Tag Result:", images)

	// push images
	result, err := push(ar.(Registrator))
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
