FROM golang:1.8

ADD . /go/src/github.com/markTward/gocloud

RUN go get -d ./...
RUN go install github.com/markTward/gocloud
