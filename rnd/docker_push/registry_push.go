package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Workflow
	Registry
}

var configFile, buildTag, event, branch, baseImage, pr *string

func init() {
	const (
		defaultConfigFile = "cicd.yaml"
		configFileUsage   = "configuration file containing project workflow values"
		buildTagUsage     = "existing image tag used as basis for further tags (required)"
		eventUsage        = "build event type from list: push, pull_request"
		branchTypeUsage   = "build branch (required)"
		prUsage           = "pull request number (required when event type is pull_request)"
		baseImageUsage    = "built image used as basis for tagging (required)"
	)
	baseImage = flag.String("image", "", baseImageUsage)
	configFile = flag.String("config", defaultConfigFile, configFileUsage)
	buildTag = flag.String("tag", "", buildTagUsage)
	event = flag.String("event", "push", eventUsage)
	branch = flag.String("branch", "", branchTypeUsage)
	pr = flag.String("pr", "", prUsage)
}

func makeTagList(repoURL string, refImage string, event string, branch string, pr string) (images []string, err error) {

	log.Println("Tagger args:", repoURL, refImage, event, branch, pr)

	// tag additional images based on build event type
	tagSep := strings.Index(refImage, ":")
	commitImage := repoURL + refImage[tagSep:]
	log.Println("commit image:", commitImage)

	images = append(images, commitImage)

	switch event {
	case "push":
		images = append(images, repoURL+":"+branch)
		if branch == "master" {
			images = append(images, repoURL+":latest")
		}
	case "pull_request":
		images = append(images, repoURL+":PR-"+pr)
	}

	log.Println("tagged images:", images)
	return images, err
}

func tagImages(src string, targets []string) (err error) {
	var stderr bytes.Buffer

	for _, target := range targets {
		cmd := exec.Command("docker", "tag", src, target)
		cmd.Stderr = &stderr
		log.Printf("docker tag from %v to %v", src, target)

		if err = cmd.Run(); err != nil {
			err = fmt.Errorf("%v", stderr.String())
			break
		}
	}

	return err
}

func validateCLInput() (err error) {

	if *baseImage == "" {
		err = fmt.Errorf("%v\n", "build image a required value; use --image option")
	}

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

	if *event == "pull_request" && *pr == "" {
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

	// TODO: return tag() to receiver and eliminate need to capture url
	switch cfg.Workflow.Registry {
	case "gcr":
		activeRegistry = &cfg.Registry.GCRRegistry
	case "docker":
		activeRegistry = &cfg.Registry.DockerRegistry
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
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}

	// make list of images to tag
	var images []string
	if images, err = makeTagList(ar.GetRepoURL(), *baseImage, *event, *branch, *pr); err != nil {
		log.Printf("error: %v", err)
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}

	// tag images
	if err = tagImages(*baseImage, images); err != nil {
		log.Printf("error: %v", err)
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}

	// push images
	var result []string
	if result, err = ar.Push(images); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	log.Println("pushed images:", result)

}
