package main

import (
	"fmt"
	"hash/maphash"
)

func main() {
	var h maphash.Hash

	h.WriteString("hello, ")
	fmt.Printf("%#x\n", h.Sum64())

	h.Write([]byte{'w', 'o', 'r', 'l', 'd'})
	fmt.Printf("%#x\n", h.Sum64())

	h.Reset()

	var h2 maphash.Hash
	h2.SetSeed(h.Seed())

	h.WriteString("same")
	h2.WriteString("same")
	fmt.Printf("%#x == %#x\n", h.Sum64(), h2.Sum64())
}
