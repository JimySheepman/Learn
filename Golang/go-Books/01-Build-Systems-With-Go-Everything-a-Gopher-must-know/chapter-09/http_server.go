package main

import (
	"fmt"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		fmt.Println(name, headers)
	}
	w.Write([]byte("Perfect!!!"))
	return
}

func main() {
	http.HandleFunc("/info", Info)
	panic(http.ListenAndServe(":8080", nil))
}
