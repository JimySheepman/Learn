package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key: ", k, "\nval: ", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello merlins")
}

func main() {

	http.HandleFunc("/", sayHelloName)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
