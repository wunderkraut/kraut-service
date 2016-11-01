package main

import (
	"time"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func main() {
	
	// Use the default handler for now
	handler := http.DefaultServeMux

	s := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}