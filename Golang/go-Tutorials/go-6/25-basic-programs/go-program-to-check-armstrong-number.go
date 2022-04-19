package main

import "fmt"

func main() {
	var number, temNumber, remainder int
	var result int = 0
	fmt.Print("Enter any three digit number: ")
	fmt.Scan(&number)

	temNumber = number

	for {
		remainder = temNumber % 10
		result += remainder * remainder * remainder
		temNumber /= 10
		if temNumber == 0 {
			break
		}
	}

	if result == number {
		fmt.Printf("%d is an Armstrong number.", number)
	} else {
		fmt.Printf("%d is not an Armstrong number.", number)
	}
}
