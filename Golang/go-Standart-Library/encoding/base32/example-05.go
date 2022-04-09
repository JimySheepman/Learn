package main

import (
	"encoding/base32"
	"os"
)

func main() {
	input := []byte("foo\x00bar")
	encoder := base32.NewEncoder(base32.StdEncoding, os.Stdout)
	encoder.Write(input)
	encoder.Close()
}
