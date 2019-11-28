FROM golang:latest

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/derekparker/delve/cmd/dlv
COPY ./ /go/src/github.com/learnCloudProvideer
WORKDIR /go/src/github.com/learnCloudProvideer
RUN apt -y update && apt -y install vim git
RUN go build main.go