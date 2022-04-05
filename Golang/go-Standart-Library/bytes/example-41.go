package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Print(string(bytes.TrimRight([]byte("453gopher8257"), "0123456789")))
}
