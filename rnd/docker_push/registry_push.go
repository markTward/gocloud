package main

// TODO: commit tag, event type and branch flags and logic
// TODO: create one each subscripts for gcr and docker

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Registry
}

const (
	defaultConfigFile  = "cicd.yaml"
	defaultRegistryURL = ""
	configFileUsage    = "configuration file containing project workflow values"
	registryURLUsage   = "provide a valid registry URL (docker and gcr currenly supported)"
)

var configFile string
var registryURL string

func init() {

	flag.StringVar(&configFile, "config", defaultConfigFile, configFileUsage)
	flag.StringVar(&configFile, "c", defaultConfigFile, configFileUsage)

	flag.StringVar(&registryURL, "url", defaultRegistryURL, registryURLUsage)
	flag.StringVar(&registryURL, "u", defaultRegistryURL, registryURLUsage)
}

func main() {

	flag.Parse()

	yamlInput, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	cfg := Config{}

	err = yaml.Unmarshal([]byte(yamlInput), &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if registryURL != "" {
		if !(strings.HasPrefix(registryURL, "gcr.io") || strings.HasPrefix(registryURL, "docker.io")) {
			fmt.Fprintf(os.Stderr, "error: unsupported registry: %v\n", strings.Split(registryURL, "/")[0])
			os.Exit(1)
		}
	} else {
		switch cfg.Registry.Name {
		case "gcr":
			registryURL = fmt.Sprintf("%v/%v/%v", cfg.Registry.Host, cfg.Registry.Project, cfg.Registry.Repo)
			// TODO: os.exec gcr_push.sh
		case "docker":
			registryURL = fmt.Sprintf("%v/%v/%v", cfg.Registry.Host, cfg.Registry.Account, cfg.Registry.Repo)
			// TODL: os.exec docker_push.sh
		default:
			fmt.Fprintf(os.Stderr, "error: unsupported registry: %v\n", cfg.Registry.Name)
			os.Exit(1)
		}
	}
	fmt.Println("Registry URL:", registryURL)

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
