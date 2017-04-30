package main

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

var configFile, buildTag, event, branch *string
var pr *int

func init() {
	const (
		defaultConfigFile = "cicd.yaml"
		configFileUsage   = "configuration file containing project workflow values (default: cicd.yaml)"
		buildTagUsage     = "existing image tag used as basis for further tags (required)"
		eventUsage        = "build event type from list: push, pull_request (default push)"
		branchTypeUsage   = "build branch (required)"
		prUsage           = "pull request number (required when event type is pr)"
	)
	configFile = flag.String("config", defaultConfigFile, configFileUsage)
	buildTag = flag.String("tag", "", buildTagUsage)
	event = flag.String("event", "push", eventUsage)
	branch = flag.String("branch", "", branchTypeUsage)
	pr = flag.Int("pr", 0, prUsage)
}

func tag(url string, tag string, event string, branch string, pr int) ([]string, error) {

	// build tag images
	var images []string

	// always tag image using git commit
	image := url + ":" + tag
	images = append(images, image)

	// tag additional images based on build event type
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

func push(r Registrator, images []string) (string, error) {
	return r.Push(images)
}

func isRegistryValid(r Registrator) bool {
	return r.IsRegistryValid()
}

func validateCLInput() (err error) {
	if *buildTag == "" {
		err = fmt.Errorf("%v\n", "error: build tag a required value; use --tag option")
	}

	if *branch == "" {
		err = fmt.Errorf("%v\n", "error: build branch a required value; use --branch option")
	}

	switch *event {
	case "push", "pull_request":
	default:
		err = fmt.Errorf("%v\n", "error: event type must be one of: push, pull_request")
	}

	if *event == "pull_request" && *pr == 0 {
		err = fmt.Errorf("%v\n", "error: event type pull_request requires a PR number; use --pr option")
	}
	return err
}

func main() {

	// parse and validate CLI
	flag.Parse()
	if err := validateCLInput(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}

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
	images, _ = tag(url, *buildTag, *event, *branch, *pr)
	fmt.Println("Tag Result:", images)

	// push images
	var result string
	if result, err = push(ar.(Registrator), images); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	fmt.Printf("Push Result: %v\n", result)

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
