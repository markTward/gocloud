package restapi

import (
	"log"
	"net/http"
)

const (
	address = ":8010"
)

func Start() {
	api := &RestAPI{
		HelloWorlder:  HelloWorldEndpoint{},
		HealthChecker: HealthCheckEndpoint{},
	}

	http.HandleFunc("/hw", api.HelloWorldHandler)
	http.HandleFunc("/healthcheck", api.HealthCheckHandler)

	log.Printf("RestAPI api: %#v %T\n", api, api)
	log.Printf("RestAPI &api: %#v %T\n", &api, &api)
	log.Printf("RestAPI *api: %#v %T\n", *api, *api)

	log.Println("listening on", address)

	log.Fatal(http.ListenAndServe(address, nil))
}
