package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	complexUrl := "https://www.example.com/path/to/?query=123&this=that#fragment"
	parsedUrl, err := url.Parse(complexUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scheme: " + parsedUrl.Scheme)
	fmt.Println("Host: " + parsedUrl.Host)
	fmt.Println("Path: " + parsedUrl.Path)
	fmt.Println("Query string: " + parsedUrl.RawQuery)
	fmt.Println("Fragment: " + parsedUrl.Fragment)

	fmt.Println("\nQuery values:")
	queryMap := parsedUrl.Query()
	fmt.Println(queryMap)

	var customURL url.URL
	customURL.Scheme = "https"
	customURL.Host = "google.com"
	newQueryValues := customURL.Query()
	newQueryValues.Set("key1", "value1")
	newQueryValues.Set("key2", "value2")
	customURL.Fragment = "bookmarkLink"
	customURL.RawQuery = newQueryValues.Encode()

	fmt.Println("\nCustom URL:")
	fmt.Println(customURL.String())
}
