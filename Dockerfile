FROM golang:1.8

ADD . /go/src/github.com/markTward/gocloud

# package dependencies
# RUN go get ./...
RUN go get golang.org/x/net/context
RUN go get google.golang.org/grpc
RUN go get github.com/spf13/cobra
RUN go get github.com/spf13/viper
RUN go get github.com/golang/mock/gomock

RUN go install github.com/markTward/gocloud
