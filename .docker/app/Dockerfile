FROM golang:1.14.6-alpine3.11
ENV PATH="$PATH:/bin/bash" 

RUN apk add --update bash curl make

WORKDIR /go/src

RUN apk add --update --upgrade curl python unzip bash gcc g++ scons 
WORKDIR /go/src
ENTRYPOINT ["tail", "-f", "/dev/null"]
