package main

import "fmt"

func main() {
	var x, n, z uint8

	x = 200
	fmt.Printf("%08b\n", x)
	n = 1
	z = x << n
	fmt.Printf("%08b\n", z)
	fmt.Printf("%d\n", z)

	x = 200
	fmt.Printf("%08b\n", x)
	n = 1
	z = x << n
	fmt.Printf("%b\n", z)
	fmt.Printf("%d\n", z)

	x = 200
	fmt.Printf("%08b\n", x)
	n = 3
	z = x << n
	fmt.Printf("%b\n", z)
	fmt.Printf("%d\n", z)
}
