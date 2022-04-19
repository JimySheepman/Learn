package main

import "fmt"

func add(x, y int) int {
	total := 0
	total = x + y
	return total
}

func main() {
	sum := add(20, 30)
	fmt.Println(sum)
}
