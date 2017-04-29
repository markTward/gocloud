package main

import (
	"fmt"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type GCR struct {
	Name    string
	Url     string
	Project string
	Repo    string
}

type Docker struct {
	Name    string
	Url     string
	Account string
	Repo    string
}

type Config struct {
	// Registry map[string]map[string]Platform
	// Registry map[string][]Platform
	Registry map[string][]interface{}
	// Registry map[string]map[string]interface{}
}

type Platform struct {
	Name    string
	Url     string
	Project string
	Repo    string
}

func main() {

	var yamlInput2 = `
registry:
  provider:
    - name: gcr
      url: gcr.io
      project: k8sdemo-159622
      repo: gocloud
    - name: docker
      url: docker.io
      account: marktward
      repo: gocloud
`
	cfg := Config{}

	err := yaml.Unmarshal([]byte(yamlInput2), &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("New Config:", cfg)
	provider := cfg.Registry["provider"]
	fmt.Printf("\n\n%#v :: %T\n", provider, provider)

	for k1, v1 := range provider {
		fmt.Printf("k1:%#v :: v1:%#v // %T :: %T\n", k1, v1, k1, v1)
		m := make(map[interface{}]interface{})
		m = v1.(map[interface{}]interface{})
		fmt.Printf("m post assert: %#v // %T\n", m, m)

		var gcrpoint interface{}
		gcrpoint = &m
		gcrassert, ok := gcrpoint.(GCR)
		fmt.Println("GCR assert ==>", gcrpoint, gcrassert, ok)

		m2 := make(map[string]string)
		m2["test"] = "ing"
		for k2, v2 := range m {
			// m2[k2] = v2
			key := k2.(string)
			value := v2.(string)
			m2[key] = value
			fmt.Printf("map k2 ==> %#v (%T) :: %#v (%T)\n", k2, k2, v2, v2)
		}

		fmt.Printf("new map: %#v (%T)\n", m2, m2)
		fmt.Println()
}
