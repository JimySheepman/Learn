package main

import "fmt"

func main() {
	a, b := 10, 2

	equal := a == b
	fmt.Printf("a == b => %v\n", equal)

	notEqual := a != b
	fmt.Printf("a != b => %v\n", notEqual)

	greater := a > b
	fmt.Printf("a > b => %v\n", greater)

	greaterOrEqual := a >= b
	fmt.Printf("a >= b => %v\n", greaterOrEqual)

	less := a < b
	fmt.Printf("a < b => %v\n", less)

	lessOrEqual := a <= b
	fmt.Printf("a <= b => %v\n", lessOrEqual)

}
