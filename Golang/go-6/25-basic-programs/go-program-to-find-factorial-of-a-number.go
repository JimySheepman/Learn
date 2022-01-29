package main

import "fmt"

var (
	factVal uint64 = 1
	i       int    = 1
	n       int
)

func factorial(n int) uint64 {
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i)
		}

	}
	return factVal
}

func main() {
	fmt.Print("Enter a positive integer between 0-50:")
	fmt.Scan(&n)
	fmt.Printf("Factorial is: %d\n", factorial(n))

}
