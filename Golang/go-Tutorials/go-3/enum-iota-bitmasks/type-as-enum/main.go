package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type HTTPMethod int

const (
	GET     HTTPMethod = 0
	POST    HTTPMethod = 1
	PUT     HTTPMethod = 2
	DELETE  HTTPMethod = 3
	PATCH   HTTPMethod = 4
	HEAD    HTTPMethod = 5
	OPTIONS HTTPMethod = 6
	TRACE   HTTPMethod = 7
	CONNECT HTTPMethod = 8
)

func handle(method HTTPMethod, headers map[string]string, uri string) {
	if method == GET {
		fmt.Println("the method is get")
	} else {
		fmt.Println("the method is not get")
	}
}

type HTTPRequest struct {
	Method  HTTPMethod        `json:"method"`
	Headers map[string]string `json:"headers"`
	Uri     string            `json:"uri"`
}

func main() {
	r := HTTPRequest{
		Method:  GET,
		Headers: map[string]string{"Accept": "application/json"},
		Uri:     "/prices",
	}

	fmt.Println(r)

	lolMethod := HTTPMethod(42)
	headers := make(map[string]string)
	handle(lolMethod, headers, "/prices")
	log.Println(GET)

	r = HTTPRequest{Method: GET, Headers: map[string]string{"Accept": "application/json"}, Uri: "/prices"}
	_, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	jsonB := []byte("{\"method\":\"GET\",\"headers\":{\"Accept\":\"application/json\"},\"uri\":\"/prices\"}")
	req := HTTPRequest{}
	err = json.Unmarshal(jsonB, &req)
	if err != nil {
		panic(err)
	}
}
