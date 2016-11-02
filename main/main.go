package main

import (
	"time"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func main() {
	
	// Use the default handler for now
	handler := http.DefaultServeMux

	server := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Add Operations Handlers
	coreHandlersAdd(server)


	// Log service ending
	log.Fatal(server.ListenAndServe())
}

// Add some initial base handlers
func coreHandlersAdd(server http.Server) error {


	return nil
}
