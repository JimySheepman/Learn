package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var requestCount int

func main() {
	http.HandleFunc("/status", status)
	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func status(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	requestCount++
	mu.Unlock()
	fmt.Fprintf(w, "OK - count : %d \n", requestCount)
}
