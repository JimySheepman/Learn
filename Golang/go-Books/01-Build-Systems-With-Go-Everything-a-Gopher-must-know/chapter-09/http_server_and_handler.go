package main

import (
	"log"
	"net/http"
)

type MyHandler struct{}

func (c *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.RequestURI {
	case "/hello":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("goodbye\n"))
	case "/goodbye":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello\n"))
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	handler := MyHandler{}
	log.Fatal(http.ListenAndServe(":8080", &handler))
}
