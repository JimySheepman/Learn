package main

import (
	"fmt"
	"log"
	"net/http"
)

var requestCount int

func main() {
	http.HandleFunc("/status", status)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func status(w http.ResponseWriter, req *http.Request) {
	requestCount++
	fmt.Fprintf(w, "OK - count : %d \n", requestCount)
}
