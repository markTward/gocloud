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

type RestAPIIniter interface {
	Init(string) error
}

func InitAPI(init RestAPIIniter, name string) error {
	return init.Init(name)
}

func (api *RestAPI) Init(name string) error {
	api.Name = name
	eps := make(Endpoints)

	ephw := Endpoint{
		id:          "hw",
		url:         "/hw",
		description: "hello world! endpoint",
	}

	epgit := Endpoint{
		id:          "git",
		url:         "/github.com/user",
		description: "github user endpoint",
	}

	eplog := Endpoint{
		id:          "login",
		url:         "/login",
		description: "user login",
	}

	eps[ephw.id] = ephw
	eps[epgit.id] = epgit
	eps[eplog.id] = eplog

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

}
