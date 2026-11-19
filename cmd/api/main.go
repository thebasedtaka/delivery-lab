package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"example.com/delivery-lab/internal/api"
)

// Later levels will inject a release version at build time.
var version = "dev"

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}
	server := &http.Server{
		Addr:              addr,
		Handler:           api.NewHandler(version),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	log.Printf("delivery-lab listening on %s", addr)
	log.Fatal(server.ListenAndServe())
}
