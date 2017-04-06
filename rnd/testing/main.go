package main

import (
	"log"
	"net/http"

	"github.com/markTward/gocloud/rnd/endpoints"
)

type RestAPI struct {
	hw endpoints.HelloWorldEndpoint
}

func (api *RestAPI) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		name = "World!"
	}

	hw := api.hw.HelloWorld(name)
	log.Printf("HelloWorld value: %s", hw)
	w.Write([]byte(hw))
}

func main() {
	api := &RestAPI{hw: endpoints.RestAPIHelloWorldEndpoint{}}

	http.HandleFunc("/hw", api.HelloWorldHandler)

	log.Println("listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
