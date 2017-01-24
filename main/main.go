package main

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
)

func main() {

	// Use the default handler for now
	handler := http.DefaultServeMux

	service := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Log server stop as an error
	log.Fatal(service.ListenAndServe())

}
