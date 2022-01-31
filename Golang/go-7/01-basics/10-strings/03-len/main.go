package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "Gopher!"
	length := len(name)
	fmt.Printf("len('%s') = %d\n", name, length)

	name = "Jimmy"
	length = len(name)
	fmt.Printf("len('%s') = %d\n", name, length)

	length = utf8.RuneCountInString(name)
	fmt.Printf("len('%s') = %d\n", name, length)
}
