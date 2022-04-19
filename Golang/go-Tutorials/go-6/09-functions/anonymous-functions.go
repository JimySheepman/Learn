package main

import "fmt"

var (
	area = func(l, b int) int {
		return l * b
	}
)

func main() {
	fmt.Println(area(20, 30))
	// Passing arguments to anonymous functions
	func(l, b int) {
		fmt.Println(l * b)
	}(20, 30)
	// Function defined to accept a parameter and return value
	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)
}
