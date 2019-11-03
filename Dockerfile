FROM golang:latest

RUN go get -u github.com/golang/dep/cmd/dep
COPY ./ /go/src/github.com/learnCloudProvideer
WORKDIR /go/src/github.com/learnCloudProvideer
RUN go build main.go
