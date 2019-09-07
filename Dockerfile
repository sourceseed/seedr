FROM golang:alpine

RUN apk add --no-cache git bash openssh zip

WORKDIR $GOPATH/src/github.com/sourceseed/seedr
COPY . .
RUN /bin/bash scripts/build.sh

