package restapi

import (
	"fmt"
	"log"
	"net/http"

	ep "github.com/markTward/gocloud/restapi/endpoints"
)

type RestAPI struct {
	ep.HealthWorlder
	ep.HealthChecker
}

func (api *RestAPI) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(api.HealthCheck())
	fmt.Fprint(w, "OK")
	log.Println(r.URL.Path, http.StatusOK)
}

func (api *RestAPI) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("r.URL.Query()[\"name\"]", r.URL.Query()["name"])

	msg, err := api.HelloWorld(r.URL.Query()["name"])
	if err != nil {
		log.Println(err)
		//TODO: write user friendly error message
		http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
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
		ep.RestAPIHelloWorldEndpoint{},
		ep.RestAPIHealthCheckEndpoint{},
	}

	http.HandleFunc("/hw", api.HelloWorldHandler)
	http.HandleFunc("/healthcheck", api.HealthCheckHandler)

	log.Println("listening on", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
