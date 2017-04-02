FROM golang:1.8

ADD . /go/src/github.com/markTward/gocloud/helloworld

# package dependencies
RUN go get golang.org/x/net/context
RUN go get google.golang.org/grpc

RUN go install github.com/markTward/gocloud/helloworld/greeter_server
RUN go install github.com/markTward/gocloud/helloworld/greeter_client
