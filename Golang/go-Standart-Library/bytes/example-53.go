package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.NewReader([]byte("Hi!")).Len())
	fmt.Println(bytes.NewReader([]byte("こんにちは!")).Len())
}
