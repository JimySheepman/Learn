package main

import "fmt"

func main() {
	var aRune rune = 'Z'
	fmt.Printf("Unicode Code point of &#39;%c&#39;: %U\n", aRune, aRune)
}
