package restapi

import (
	"log"
	"net/http"

	"github.com/markTward/gocloud/restapi/handlers"
)

const (
	address = ":8010"
)

func Start() {
	//TODO: handle 404 and other http errors
	http.HandleFunc("/hw", handlers.HelloWorld)
	http.HandleFunc("/healthcheck", handlers.HealthCheck)
	log.Fatal(http.ListenAndServe(address, nil))
}
