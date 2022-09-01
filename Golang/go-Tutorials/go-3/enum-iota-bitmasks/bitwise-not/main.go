package main

import "fmt"

func main() {
	var x, z uint8
	x = 200
	fmt.Printf("%08b\n", x)

	z = ^x

	fmt.Printf("%08b\n", z)
	fmt.Printf("%d\n", z)
}
