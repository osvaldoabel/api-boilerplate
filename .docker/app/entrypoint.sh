#!/bin/bash

chmod -R 777 storage
go mod tidy  && 
go test ./... && 
go run api/server.go