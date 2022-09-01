package main

import "fmt"

func main() {
	var x uint8

	x = 1
	fmt.Printf("%08b\n", x)

	x = 2
	fmt.Printf("%08b\n", x)

	x = 255
	fmt.Printf("%08b\n", x)
}
