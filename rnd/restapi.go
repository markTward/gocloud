package main

import "fmt"

type RestAPI struct {
	Name string
	Endpoints
}

type RestAPIGreediMonki RestAPI

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

type LoginEndpoint struct{}

type RestAPIInit interface {
	init(string) RestAPI
}

func (api *RestAPI) init(name string) {
	api.Name = name
	eps := make(Endpoints)

	ephw := HelloWorldEndpoint{}
	ephw.ep.id = "hw"
	ephw.ep.url = "/hw"
	ephw.ep.description = "hello world! endpoint"

	epgit := GitEndpoint{
		Endpoint{
			id:          "git",
			url:         "/github.com/user",
			description: "github user endpoint",
		},
	}

	eps[ephw.ep.id] = ephw.ep
	eps[epgit.id] = epgit.Endpoint

	api.Endpoints = eps

}

func (api *RestAPIGreediMonki) init(name string) {
	api.Name = name
	eps := make(Endpoints)

	eps["login"] = Endpoint{
		id:          "login",
		url:         "/login",
		description: "user login",
	}

	api.Endpoints = eps

}

func main() {

	api := &RestAPI{}
	api.init("GoCloud")

	fmt.Println("RestAPI:", api.Name)
	for _, ep := range api.Endpoints {
		fmt.Println("\tEndpoint:", ep)
	}

	apigm := &RestAPIGreediMonki{}
	apigm.init("GreediMonki")

	fmt.Println("RestAPI:", apigm.Name)
	for _, ep := range apigm.Endpoints {
		fmt.Println("\tEndpoint:", ep)
	}

}
