package main

import (
	"github.com/gichohi/go-rest.git/routes"
	"log"
	"net/http"
)

const (
	SERVICE_PORT 	= ":8082"
)

func main() {

	http.Handle("/", routes.Handlers())
	log.Fatal(http.ListenAndServe(SERVICE_PORT, nil))
}