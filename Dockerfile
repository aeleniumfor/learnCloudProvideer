FROM sasenomura/golang-nvim-dev:latest

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/derekparker/delve/cmd/dlv
COPY ./ /go/src/github.com/learnCloudProvideer
WORKDIR /go/src/github.com/learnCloudProvideer
RUN go build main.go