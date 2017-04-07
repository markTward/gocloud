package restapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/markTward/gocloud/restapi/endpoints"
)

type RestAPI struct {
	hw endpoints.HelloWorldEndpoint
	hc endpoints.HealthCheckEndpoint
}

func (api *RestAPI) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(api.hc.HealthCheck())
	fmt.Fprint(w, "OK")
	log.Println(r.URL.Path, http.StatusOK)
}

func (api *RestAPI) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("r.URL.Query()[\"name\"]", r.URL.Query()["name"])
	msg, err := api.hw.HelloWorld(r.URL.Query()["name"])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("%s?%s; Message:%s", r.URL.Path, r.URL.RawQuery, msg)
	fmt.Fprint(w, msg)
}

const (
	address = ":8010"
)

func Start() {
	api := &RestAPI{
		hw: endpoints.RestAPIHelloWorldEndpoint{},
		hc: endpoints.RestAPIHealthCheckEndpoint{},
	}

	http.HandleFunc("/hw", api.HelloWorldHandler)
	http.HandleFunc("/healthcheck", api.HealthCheckHandler)

	log.Println("listening on", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
