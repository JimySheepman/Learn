package main

import "fmt"

func hello() {
	fmt.Println("Hello, World!")
}

func UTF_8() {
	for i := 60; i < 122; i++ {
		// ascii code, binary, decimal, char
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}

func number() {
	num := 42

	// decimal
	fmt.Println(num)
	// binary
	fmt.Printf("%d - %b\n", num, num)
	// hexadecimal
	fmt.Printf("%d \t %b \t %#X \n", num, num, num)
	// loop
	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}

func main() {
	hello()
	UTF_8()
	number()
}
