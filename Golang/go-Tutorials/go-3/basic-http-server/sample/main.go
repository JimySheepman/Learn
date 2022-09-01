package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	abTestingServer := http.Server{
		Addr:         "127.0.0.1:9899",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      &MyHandler{},
	}
	log.Fatal(abTestingServer.ListenAndServe())
}

type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO
}
