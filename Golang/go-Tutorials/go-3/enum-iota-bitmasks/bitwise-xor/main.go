package main

import "fmt"

func main() {
	var x, y, z uint8
	x = 200
	fmt.Printf("%08b\n", x)

	y = 100
	fmt.Printf("%08b\n", y)

	z = x ^ y

	fmt.Printf("%08b\n", z)

	fmt.Printf("%d\n", z)
}
