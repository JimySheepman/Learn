package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	var needle []byte
	var haystack [][]byte

	i := sort.Search(len(haystack), func(i int) bool {
		return bytes.Compare(haystack[i], needle) >= 0
	})
	if i < len(haystack) && bytes.Equal(haystack[i], needle) {
		fmt.Println("Found it!")
	}
}
