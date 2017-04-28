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
	// Registry map[string][]interface{}
	Registry map[string]map[string]interface{}
}

type Platform struct {
	Name    string
	Url     string
	Project string
	Repo    string
}

func main() {

	// p11 := GCRPlatform{
	// 	Url:     "grc.io",
	// 	Project: "k8sdemo-159622",
	// 	Repo:    "gocloud",
	// 	Name:    "gcr",
	// }
	//
	// p22 := DockerPlatform{
	// 	Name:    "docker",
	// 	Url:     "docker.io",
	// 	Account: "markTward",
	// 	Repo:    "gocloud",
	// }
	//
	// var pz2 []interface{}
	// pz2 = append(pz2, p11, p22)
	// fmt.Printf("%#v\n", pz2)
	//
	// for i, v := range pz2 {
	// 	fmt.Printf("%v:  %#v (%T)\n", i, v, v)
	//
	// 	switch r := reflect.TypeOf(v).String(); r {
	// 	case "main.GCRPlatform":
	// 		fmt.Println("GCR")
	// 	case "main.DockerPlatform":
	// 		fmt.Println("Docker")
	// 	default:
	// 		fmt.Println("Unknown")
	// 	}
	//
	// }

	// 	var yamlInput1 = `
	//   registry:
	//     provider:
	//       gcrplatform:
	//         name: gcr
	//         url: gcr.io
	//         project: k8sdemo-159622
	//         repo: gocloud
	//       dockerplatform:
	//         name: docker
	//         url: docker.io
	//         project: markTward
	//         repo: gocloud
	// `
	// 	var yamlInput2 = `
	// registry:
	//   provider:
	//     - name: gcr
	//       url: gcr.io
	//       project: k8sdemo-159622
	//       repo: gocloud
	//     - name: docker
	//       url: docker.io
	//       account: marktward
	//       repo: gocloud

	var yamlInput3 = `
  registry:
    provider:
      gcr:
        name: gcr
        url: gcr.io
        project: k8sdemo-159622
        repo: gocloud
      docker:
        name: docker
        url: docker.io
        account: markTward
        repo: gocloud
  `
	cfg := Config{}

	err := yaml.Unmarshal([]byte(yamlInput3), &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("New Config:", cfg)
	provider := cfg.Registry["provider"]
	fmt.Printf("\n\n%#v :: %T\n", provider, provider)

	r1, ok := cfg.Registry["provider"].(map[string]interface{})
	// r1, ok := cfg.Registry["provider"]["gcr"].(map[interface{}]interface{})
	if ok {
		fmt.Println("R1 OK:", r1)
	} else {
		fmt.Println("r1 NOT ok:")
	}
}
