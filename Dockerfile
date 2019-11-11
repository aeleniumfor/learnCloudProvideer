FROM golang:latest

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/derekparker/delve/cmd/dlv
COPY ./ /go/src/github.com/learnCloudProvideer
WORKDIR /go/src/github.com/learnCloudProvideer
RUN apt -y update && apt -y install vim git
RUN go build main.go

RUN mkdir -p ~/.vim/autoload ~/.vim/bundle && \
    curl -LSso ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim && \
    echo 'execute pathogen#infect()' >> ~/.vimrc && \
    echo 'syntax on' >> ~/.vimrc && \
    echo 'filetype plugin indent on' >> ~/.vimrc && \
    git clone https://github.com/fatih/vim-go.git ~/.vim/bundle/vim-go