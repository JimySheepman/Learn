package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	runes := []rune{
		-1,
		0x110000,
		utf8.RuneError,
	}
	for i, c := range runes {
		buf := make([]byte, 3)
		size := utf8.EncodeRune(buf, c)
		fmt.Printf("%d: %d %[2]s %d\n", i, buf, size)
	}
}
