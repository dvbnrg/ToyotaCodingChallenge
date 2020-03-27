package main

import (
	"github.com/dvbnrg/ToyotaCodingChallenge/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":80", router))
}
