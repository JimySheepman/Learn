package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func invalid_code() {
	a := "stored ina a"
	// b := "stored ina b"
	// b is not beig used - invalid code
	fmt.Println("a- ", a)
}

func http_get_example() {
	// wit error checking
	res, err := http.Get("http://www.geekwiseacademy.com/")
	if err != nil {
		log.Fatal()
	}

	page, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}

func main() {
	invalid_code()
	http_get_example()
}
