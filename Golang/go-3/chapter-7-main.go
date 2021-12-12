package main

import "fmt"

func main() {
	n := 2548
	n2 := 0x9F4
	n3 := 02454
	n4 := 1324
	n5 := 0x9F4
	b := make([]byte, 0)
	b = append(b, 255)
	b = append(b, 10)
	str := `spring rain:
browsing under an umbrella
at the picture-book store`
	fmt.Printf("%X\n", n)
	fmt.Printf("%x\n", n2)
	fmt.Printf("Decimal : %d \n", n2)
	fmt.Printf("decimal : %d\n", n3)
	fmt.Printf("octal: %o\n", n4)
	fmt.Printf("octal with prefix: %O\n", n4)
	fmt.Println(b)
	fmt.Printf("Decimal: %d\n", n5)
	fmt.Printf("Binary : %b\n", n5)
	fmt.Println(str)
	interpreted := "I love spring"
	fmt.Println(interpreted)
	s := "I Love Golang"
	for _, v := range s {
		fmt.Printf("Unicode code point : %U - character '%c' - binary %b - hex %X - Decimal %d\n ", v, v, v, v, v)
	}
	var aRune rune = 'Z'
	fmt.Printf("Unicode Code point of &#39;%c&#39;: %U\n", aRune, aRune)
}
