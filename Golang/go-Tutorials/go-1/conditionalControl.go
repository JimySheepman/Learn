package main

import "fmt"

func main() {
	num1 := 5
	num2 := 3

	if num1 < num2 {
		fmt.Printf("if condition is true")
	} else if num1 > num2 {
		fmt.Printf("else if condition is true")
	} else {
		fmt.Printf("nothing above is true, execute else fallback")
	}

	a := 10
	if a < 20 {
		fmt.Println("\nouter if")
		if a >= 10 {
			fmt.Println("inner if")
		}
	}

	grade := 'B'
	switch grade {
	case 'A':
		fmt.Println("Excellent!")
	case 'B':
		fmt.Println("Very good")
	case 'C':
		fmt.Println("Well done")
	case 'D':
		fmt.Println("You passed")
	case 'F':
		fmt.Println("You failed")
	default:
		fmt.Println("Invalid grade")
	}
	fmt.Println("Your grade is", grade)
}
