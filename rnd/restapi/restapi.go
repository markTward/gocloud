package main

import (
	"fmt"
	"os"
)

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

const (
	// TODO: addressDB must match k8s charts and template values
	addressDB          = "gocloud-grpc:8000"
	defaultServiceName = "gocloud-grpc"
	defaultServicePort = "8888"
	defaultName        = "World"
	timeout            = 1
)

func getServiceAddress() string {
	// svc := os.Getenv("GRPC_HW_SERVICE_NAME")
	// port := os.Getenv("GRPC_HW_SERVICE_PORT")
	var svc, port string

	if svc = os.Getenv("GRPC_HW_SERVICE_NAME"); svc == "" {
		svc = defaultServiceName
	}
	if port = os.Getenv("GRPC_HW_SERVICE_PORT"); port == "" {
		port = defaultServicePort
	}
	return fmt.Sprintf("%s:%s\n", svc, port)
}

func main() {

	api := &RestAPI{}
	InitAPI(api, "GoCloud")

	fmt.Println("RestAPI:", api.Name)
	for _, ep := range api.Endpoints {
		fmt.Println("\tEndpoint:", ep)
	}

	serviceAddress := getServiceAddress()
	fmt.Println("serviceAddress =", serviceAddress)

}
