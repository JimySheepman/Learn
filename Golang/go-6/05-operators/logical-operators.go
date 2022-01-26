package main

import "fmt"

func main() {
	var x, y, z = 10, 20, 30

	fmt.Println(x < y && x > z)
	fmt.Println(x < y || x > z)
	fmt.Println(x == y && x > z)
}
