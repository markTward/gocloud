package main

import "fmt"

type RestAPI struct {
	eps Endpoints
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
    Endpoint {
		id:          "git",
		url:         "/github.com/user",
		description: "githug user endpoint",
  }
	}

	eps := make(Endpoints)
	eps[ephw.id] = ephw.Endpoint

	api := RestAPI{eps: eps}

	fmt.Println(api.eps, eps, ephw, epgit)
	// fmt.Printf("%v / %T\n", x, x)
	fmt.Printf("%v / %T\n", ephw, ephw)

}
