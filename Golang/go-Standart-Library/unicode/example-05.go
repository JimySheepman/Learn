package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%t\n", unicode.IsNumber('Ⅷ'))
	fmt.Printf("%t\n", unicode.IsNumber('A'))
}
