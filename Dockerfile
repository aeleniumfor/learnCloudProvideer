FROM sasenomura/golang-nvim-dev:v1.1

RUN go get -u github.com/golang/dep/cmd/dep && go get -u github.com/derekparker/delve/cmd/dlv
COPY ./ /go/src/github.com/learnCloudProvideer
WORKDIR /go/src/github.com/learnCloudProvideer