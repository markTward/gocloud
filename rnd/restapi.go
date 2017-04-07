package main

import "fmt"

type RestAPI struct {
	Name string
	Endpoints
}

type RestAPIMPC RestAPI
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

type RestAPIIniter interface {
	Init(string) error
}

func InitAPI(init RestAPIIniter, name string) error {
	return init.Init(name)
}

func (api *RestAPI) Init(name string) error {
	// api.Name = name
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
	return nil
}

func (api *RestAPIGreediMonki) Init(name string) error {
	// api.Name = name
	eps := make(Endpoints)

	eps["login"] = Endpoint{
		id:          "login",
		url:         "/login",
		description: "user login",
	}

	api.Endpoints = eps
	return nil
}

func main() {

	api := &RestAPI{}
	InitAPI(api, "GoCloud")

	fmt.Println("RestAPI:", api.Name)
	for _, ep := range api.Endpoints {
		fmt.Println("\tEndpoint:", ep)
	}

	apigm := &RestAPIGreediMonki{}
	InitAPI(apigm, "GreediMonki")

	fmt.Println("RestAPI:", apigm.Name)
	for _, ep := range apigm.Endpoints {
		fmt.Println("\tEndpoint:", ep)
	}

}
