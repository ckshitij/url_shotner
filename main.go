package main

import (
	"fmt"
	"log"

	"github.io/ckshitij/url-shortner/config"
	"github.io/ckshitij/url-shortner/server"
)

func main() {

	cfg := config.LoadServiceConfig()
	log.Printf("starting server at %s ....", fmt.Sprintf(":%s", cfg.Server.Port))
	if err := server.NewHTTPServer(cfg).ListenAndServe(); err != nil {
		log.Fatal("failed to start a server with error ", err)
	}
}
