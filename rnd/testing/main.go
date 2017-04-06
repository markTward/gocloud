package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/markTward/gocloud/rnd/testing/endpoints"
)

type RestAPI struct {
	hw endpoints.HelloWorldEndpoint
}

func (api *RestAPI) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("r.URL.Query()[\"name\"]", r.URL.Query()["name"])
	msg, err := api.hw.HelloWorld(r.URL.Query()["name"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("%s?%s; Message:%s", r.URL.Path, r.URL.RawQuery, msg)
	fmt.Fprintf(w, "%s", msg)
}

func main() {
	api := &RestAPI{hw: endpoints.RestAPIHelloWorldEndpoint{}}

	http.HandleFunc("/hw", api.HelloWorldHandler)

	log.Println("listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
