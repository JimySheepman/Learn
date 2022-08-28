package main

import "fmt"

func main() {
	n := 0x9F4
	fmt.Printf("decimal : %d\n", n)

	n2 := 02454
	fmt.Printf("decimal : %d\n", n2)

	n3 := 1324
	fmt.Printf("octal : %o\n", n3)

	fmt.Printf("octal with prefix : %O\n", n3)
}
