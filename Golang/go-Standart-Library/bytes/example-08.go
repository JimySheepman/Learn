package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.EqualFold([]byte("Go"), []byte("Go")))
	fmt.Println(bytes.EqualFold([]byte("Go"), []byte("gO")))
	fmt.Println(bytes.EqualFold([]byte("Go"), []byte("go")))
}
