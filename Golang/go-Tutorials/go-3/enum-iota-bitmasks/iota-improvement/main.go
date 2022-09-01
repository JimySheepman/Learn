package main

import "fmt"

type HTTPMethod int

const (
	GET HTTPMethod = iota
	POST
	PUT
	DELETE
	PATCH
	HEAD
	OPTIONS
	TRACE
	CONNECT
)

func main() {
	fmt.Println(PUT)
}
