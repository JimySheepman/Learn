package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	fmt.Printf("Fields are: %q", bytes.FieldsFunc([]byte("  foo1;var2,ba34..."), f))
}
