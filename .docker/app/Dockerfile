FROM golang:1.14.6-alpine3.11
ENV PATH="$PATH:/bin/bash" 

RUN apk add --update bash curl make

WORKDIR /go/src

RUN apk add --update --upgrade curl python unzip bash gcc g++ scons 

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

WORKDIR /go/src
