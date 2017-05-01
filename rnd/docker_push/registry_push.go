package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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

func tagImages(src string, targets []string) (err error) {
	var stderr bytes.Buffer

	for _, target := range targets {
		cmd := exec.Command("docker", "tag", src, target)
		cmd.Stderr = &stderr
		log.Printf("attempt docker tag from %v to %v", src, target)

		if err = cmd.Run(); err != nil {
			err = fmt.Errorf("%v", stderr.String())
			break
		}
	}

	return err
}

func tag(url string, tag string, event string, branch string, pr int) (images []string, err error) {

	// source tag image using docker build git commit tag
	image := url + ":" + tag

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

	// tag additional target images
	if err = tagImages(image, images); err == nil {
		// add source image to list
		images = append(images, image)
	}
	return images, err
}

func push(r Registrator, images []string) ([]string, error) {
	return r.Push(images)
}

func validateCLInput() (err error) {
	if *buildTag == "" {
		err = fmt.Errorf("%v\n", "build tag a required value; use --tag option")
	}

	if *branch == "" {
		err = fmt.Errorf("%v\n", "build branch a required value; use --branch option")
	}

	switch *event {
	case "push", "pull_request":
	default:
		err = fmt.Errorf("%v\n", "event type must be one of: push, pull_request")
	}

	if *event == "pull_request" && *pr == 0 {
		err = fmt.Errorf("%v\n", "event type pull_request requires a PR number; use --pr option")
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
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
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

	// assert activeRegistry as type Registrator to access methods
	ar := activeRegistry.(Registrator)

	// validate registry has required values
	if err = ar.IsRegistryValid(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// authenticate credentials for registry
	if err = ar.Authenticate(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// tag images
	var images []string
	if images, err = tag(url, *buildTag, *event, *branch, *pr); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	log.Println("tagged images:", images)

	// push images
	var result []string
	if result, err = push(ar, images); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	log.Println("pushed images:", result)

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
