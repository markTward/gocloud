package main

import "fmt"

type RestAPI struct {
	Name string
	Endpoints
}

type Endpoints map[string]Endpoint

type Endpoint struct {
	id          string
	url         string
	description string
}

type HelloWorldEndpoint struct {
	Endpoint
}

type GitEndpoint struct {
	Endpoint
}

func main() {

	ephw := HelloWorldEndpoint{}
	ephw.id = "hw"
	ephw.url = "/hw"
	ephw.description = "hello world! endpoint"

	epgit := GitEndpoint{
		Endpoint{
			id:          "git",
			url:         "/github.com/user",
			description: "githug user endpoint",
		},
	}

	eps := make(Endpoints)
	eps[ephw.id] = ephw.Endpoint
	eps[epgit.id] = epgit.Endpoint

	api := RestAPI{Endpoints: eps}

	fmt.Println("RestAPI:", api.Name)

	for _, ep := range api.Endpoints {
		fmt.Println("\tEndpoint:", ep)
	}

}
