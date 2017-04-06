package endpoints

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/markTward/gocloud/proto"
	"google.golang.org/grpc"
)

const (
	addressDB   = "gocloud-grpc:8000"
	defaultName = "World!"
	timeout     = 1
)

type HelloWorldEndpoint interface {
	HelloWorld([]string) (string, error)
}

type RestAPIHelloWorldEndpoint struct{}

func (api RestAPIHelloWorldEndpoint) HelloWorld(names []string) (string, error) {
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
		msg, grpcErr = "", fmt.Errorf("grpc: server unreachable, %s", err)
		// return "", fmt.Errorf("grpc: server unreachable, %s", err)
	} else {
		msg, grpcErr = rpc.Message, nil
		// fmt.Fprint(w, rpc.Message)
	}

	return msg, grpcErr
}
