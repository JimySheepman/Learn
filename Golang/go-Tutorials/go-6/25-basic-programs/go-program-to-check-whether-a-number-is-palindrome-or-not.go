package main

import "fmt"

func main() {
	var number, remainder, temp int
	var reverse int = 0

	fmt.Print("Enter any positive integer : ")
	fmt.Scan(&number)

	temp = number

	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10

		if number == 0 {
			break
		}
	}

	if temp == reverse {
		fmt.Printf("%d is a Palidrome\n", temp)
	} else {
		fmt.Printf("%d is not a Palidrome\n", temp)
	}
}
