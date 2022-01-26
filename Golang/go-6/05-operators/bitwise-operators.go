package main

import "fmt"

func main() {
	var x uint = 9  // -> 0000 1001
	var y uint = 65 // -> 0100 0001
	var z uint

	z = x & y
	fmt.Println("x & y =", z)

	z = x | y
	fmt.Println("x | y =", z)

	z = x ^ y
	fmt.Println("x ^ y =", z)

	z = x << 1
	fmt.Println("x << 1 =", z)

	z = x >> 1
	fmt.Println("x >> 1 =", z)
}
