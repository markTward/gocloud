package main

import (
	"fmt"
	"reflect"
)

type GCRPlatform struct {
	Name    string
	Url     string
	Project string
	Repo    string
}

type DockerPlatform struct {
	Name    string
	Url     string
	Account string
	Repo    string
}

func main() {

	p11 := GCRPlatform{
		Url:     "grc.io",
		Project: "k8sdemo-159622",
		Repo:    "gocloud",
		Name:    "gcr",
	}

	p22 := DockerPlatform{
		Name:    "docker",
		Url:     "docker.io",
		Account: "markTward",
		Repo:    "gocloud",
	}

	var pz2 []interface{}
	pz2 = append(pz2, p11, p22)
	fmt.Printf("%#v\n", pz2)

	for i, v := range pz2 {
		fmt.Printf("%v:  %#v (%T)\n", i, v, v)
		r := reflect.TypeOf(v)
		fmt.Printf("Platform Type: %v\n", r)

	}

}
