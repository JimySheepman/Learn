package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/homepage", homepageHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
func homepageHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome to my homepage")
	fmt.Fprintln(writer, "I am max")

}
