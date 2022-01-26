package main

import "fmt"

func sum(x, y int) int {
	fmt.Println("1")
	return x + y
}

func partialSum(x int) func(int) int {
	fmt.Println("2")
	return func(y int) int {
		fmt.Println("3")
		return sum(x, y)
	}
}

func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func main() {
	parial := partialSum(3)
	fmt.Println(parial(7))
	// 5*5 + 6*6 + 7*7
	fmt.Println(squareSum(5)(6)(7))
}
