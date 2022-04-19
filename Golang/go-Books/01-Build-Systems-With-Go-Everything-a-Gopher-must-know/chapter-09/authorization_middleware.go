package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

type MyHandler struct{}

func (mg *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Perfect!!!"))
	return
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		autType := strings.Split(header, " ")
		fmt.Println(autType)
		if len(autType) != 2 || autType[0] != "Basic" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		credentials, err := base64.StdEncoding.DecodeString(autType[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if string(credentials) == "Open Sesame" {
			next.ServeHTTP(w, r)
		}
	})
}

func main() {
	targetHandler := MyHandler{}
	panic(http.ListenAndServe(":8080", AuthMiddleware(&targetHandler)))
}
