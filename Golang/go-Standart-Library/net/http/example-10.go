package main

import (
	"fmt"
	"log"
	"net/http"
)

func newPeopleHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the people handler.")
	})
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/resources", http.NotFoundHandler())
	mux.Handle("/resources/people/", newPeopleHandler())

	log.Fatal(http.ListenAndServe(":8080", mux))
}
