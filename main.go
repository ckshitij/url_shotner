package main

import (
	"log"

	"github.io/ckshitij/url-shortner/server"
)

func main() {

	log.Println("starting server at port :8080...")
	if err := server.NewHTTPServer().ListenAndServe(); err != nil {
		log.Fatal("failed to start a server")
	}
}
