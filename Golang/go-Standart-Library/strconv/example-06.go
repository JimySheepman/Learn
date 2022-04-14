package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("rune (ascii):")
	b = strconv.AppendQuoteRuneToASCII(b, '☺')
	fmt.Println(string(b))

}
