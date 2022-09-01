package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40}
	b := []int{1, 1, 1, 1, 1}

	copy(b, a)
	fmt.Printf("a : %v - b : %v\n", a, b)

	a = []int{10, 20, 30, 40}
	b = []int{1, 1}
	copy(b, a)
	fmt.Printf("a: %v - b : %v\n", a, b)

	a = []int{10, 20, 30, 40}
	b = make([]int, 4)
	copy(b, a)
	fmt.Printf("a: %v - b : %v\n", a, b)
}
