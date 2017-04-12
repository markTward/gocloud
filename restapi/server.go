package restapi

import (
	"log"
	"net/http"

	ep "github.com/markTward/gocloud/restapi/endpoints"
)

const (
	address = ":8010"
)

func Start() {
	api := &ep.RestAPI{
		HelloWorlder:  ep.HelloWorldEndpoint{},
		HealthChecker: ep.HealthCheckEndpoint{},
	}

	http.HandleFunc("/hw", api.HelloWorldHandler)
	http.HandleFunc("/healthcheck", api.HealthCheckHandler)

	log.Printf("RestAPI api: %#v %T\n", api, api)
	log.Printf("RestAPI &api: %#v %T\n", &api, &api)
	log.Printf("RestAPI *api: %#v %T\n", *api, *api)

	log.Println("listening on", address)

	log.Fatal(http.ListenAndServe(address, nil))
}
