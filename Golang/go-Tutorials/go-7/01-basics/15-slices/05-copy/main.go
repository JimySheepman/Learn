package main

import "fmt"

func main() {
	one := []int{1, 2}
	other := []int{3, 4}

	fmt.Printf("one: %v\n", one)
	fmt.Printf("other: %v\n", other)

	all := copy(one, other)

	fmt.Printf("oll: %v\n", all)
	fmt.Printf("one: %v\n", one)
	fmt.Printf("other: %v\n", other)
}
