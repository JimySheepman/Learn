package main

import "fmt"

func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return
}

func main() {
	var a, p int
	a, p = rectangle(20, 30)
	fmt.Println("Area: ", a)
	fmt.Println("Parameter: ", p)
}
