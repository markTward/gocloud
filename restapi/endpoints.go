package restapi

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	pb "github.com/markTward/gocloud/proto"
	"google.golang.org/grpc"
)

const (
	// TODO: addressDB must match k8s charts and template values
	addressDB   = "gocloud-grpc:8000"
	defaultName = "World"
	timeout     = 1
)

type RestAPI struct {
	HelloWorlder
	HealthChecker
}

type HelloWorlder interface {
	HelloWorld([]string) (string, error)
}

type HealthChecker interface {
	HealthCheck() int
}

// resources for servicing hello world endpoint
type HelloWorldEndpoint struct{}

func (api HelloWorldEndpoint) HelloWorld(names []string) (string, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addressDB, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(timeout*time.Second))
	if err != nil {
		return "", fmt.Errorf("grpc: server unreachable, %s", err)
	}
	defer conn.Close()

	// Greeter Client
	c := pb.NewGreeterClient(conn)

	// handle 0-to-Many qs names
	name := defaultName
	if len(names) != 0 {
		name = strings.Join(names, ", ")
	}

	// grpc attempt
	var msg string
	var grpcErr error

	rpc, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Printf("%v", err)
		msg, grpcErr = "", fmt.Errorf("%s", err)
	} else {
		msg, grpcErr = rpc.Message, nil
	}

	return msg, grpcErr
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

// resources for servicing health check endpoint
type HealthCheckEndpoint struct{}

func (api HealthCheckEndpoint) HealthCheck() int {
	return http.StatusOK
}

func (api *RestAPI) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(api.HealthCheck())
	fmt.Fprint(w, "OK")
	log.Println(r.URL.Path, http.StatusOK)
}
