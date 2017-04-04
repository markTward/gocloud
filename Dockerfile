FROM golang:1.8

ADD . /go/src/github.com/markTward/gocloud

RUN go get -v -d ./...
RUN go install -v github.com/markTward/gocloud
