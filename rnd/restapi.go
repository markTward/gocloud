package main

import "fmt"

type RestAPI struct {
	Name string
	Endpoints
	RestAPIIniter
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

func InitAPI(init RestAPIIniter) error {
	return init.Init(string)
}

func (api *RestAPI) Init() error {
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
	return nil
}

func (api *RestAPIGreediMonki) Init() error {
	api.Name = name
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

	api := &RestAPI{Name: "GoCloud"}
	InitAPI(api)

	fmt.Println("RestAPI:", api.Name)
	for _, ep := range api.Endpoints {
		fmt.Println("\tEndpoint:", ep)
	}

	apigm := &RestAPIGreediMonki{Name: "GreediMonki"}
	Init(apigm)

	fmt.Println("RestAPI:", apigm.Name)
	for _, ep := range apigm.Endpoints {
		fmt.Println("\tEndpoint:", ep)
	}

}
