package main

import (
	"fmt"
	"log"
	"net/http"
	app "osvaldoabel/users-api/src"
)

var (
	port    int
	baseURL string
)

type Headers map[string]string

func init() {
	port = 80
	baseURL = fmt.Sprintf("http://localhost:%d", port)
}

func main() {

	app.Start()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
