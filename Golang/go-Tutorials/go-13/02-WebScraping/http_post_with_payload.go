package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {

	response, err := http.PostForm(
		"http://example.com/form",
		url.Values{
			"username": {"MyUsername"},
			"password": {"123"},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	log.Println(response.Header)
}
