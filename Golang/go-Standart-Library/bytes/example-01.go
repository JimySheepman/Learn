package main

import (
	"bytes"
	"fmt"
)

func main() {
	var a, b []byte
	if bytes.Compare(a, b) < 0 {
		fmt.Println("a less b")
	}

	if bytes.Compare(a, b) <= 0 {
		fmt.Println("a less or equal b")
	}

	if bytes.Compare(a, b) > 0 {
		fmt.Println("a greater b")
	}

	if bytes.Compare(a, b) >= 0 {
		fmt.Println("a greater or equal b")
	}

	if bytes.Equal(a, b) {
		fmt.Println("a equal b")
	}

	if !bytes.Equal(a, b) {
		fmt.Println("a not equal b")
	}
}
