package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	myServer := &http.Server{
		Addr:         "127.0.0.1:8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(myServer.ListenAndServe())
}
