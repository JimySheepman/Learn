package main

import (
	"bytes"
	"fmt"
)

func main() {
	rs := bytes.Runes([]byte("go gopher"))
	for _, r := range rs {
		fmt.Printf("%#U\n", r)
	}
}
