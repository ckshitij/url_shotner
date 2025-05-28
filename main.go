package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.io/ckshitij/url-shortner/config"
	"github.io/ckshitij/url-shortner/server"
)

func handleCleanup(sigChan chan os.Signal, server *http.Server, wg *sync.WaitGroup) {
	<-sigChan
	log.Println("shutting down server gracefully...")
	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Println("error while server shutdown")
	}
	log.Println("shutting down done.")
	close(sigChan)
	wg.Done()
}

func startServer(server *http.Server, cfg *config.ServiceConfig, wg *sync.WaitGroup) {
	log.Printf("starting server at %s ....", fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))
	if err := server.ListenAndServe(); err != nil {
		log.Println("server error ", err)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	cfg := config.LoadServiceConfig()
	server := server.NewHTTPServer(cfg)

	wg.Add(2)
	notify := make(chan os.Signal, 1)
	signal.Notify(notify, os.Interrupt)
	go handleCleanup(notify, server, &wg)
	go startServer(server, cfg, &wg)

	wg.Wait()
}
