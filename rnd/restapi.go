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
	ep Endpoint
}

type GitEndpoint struct {
	Endpoint
}

func initAPI(n string) RestAPI {
	api := RestAPI{Name: n}
	eps := make(Endpoints)

	ephw := HelloWorldEndpoint{}
	ephw.ep.id = "hw"
	ephw.ep.url = "/hw"
	ephw.ep.description = "hello world! endpoint"

	epgit := GitEndpoint{
		Endpoint{
			id:          "git",
			url:         "/github.com/user",
			description: "githug user endpoint",
		},
	}

	eps[ephw.ep.id] = ephw.ep
	eps[epgit.id] = epgit.Endpoint

	api.Endpoints = eps

	return api
}

func main() {
	// eps := make(Endpoints)
	//
	// api := RestAPI{Endpoints: eps}
	//
	// ephw := HelloWorldEndpoint{}
	// ephw.ep.id = "hw"
	// ephw.ep.url = "/hw"
	// ephw.ep.description = "hello world! endpoint"
	//
	// epgit := GitEndpoint{
	// 	Endpoint{
	// 		id:          "git",
	// 		url:         "/github.com/user",
	// 		description: "githug user endpoint",
	// 	},
	// }
	//
	// eps[ephw.ep.id] = ephw.ep
	// eps[epgit.id] = epgit.Endpoint
	api := initAPI("GreediMonki")
	fmt.Println("RestAPI:", api.Name)

	for _, ep := range api.Endpoints {
		fmt.Println("\tEndpoint:", ep)
	}

}
