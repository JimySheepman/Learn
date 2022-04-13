package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.ServeMux{}
	mux.HandleFunc("/sendstrailers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Trailer", "AtEnd1,AtEnd2")
		w.Header().Add("Trailer", "AtEnd3")

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		w.Header().Set("AtEnd1", "value 1")
		io.WriteString(w, "This HTTP response has both headers before this text and trailers at the end.\n")
		w.Header().Set("AtEnd2", "value 2")
		w.Header().Set("AtEnd3", "value 3")
	})
	log.Fatal(http.ListenAndServe(":8080", &mux))
}
