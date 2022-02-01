package main

import "fmt"

func main() {
	if 1 == 1 {
		fmt.Println("1 is equal to 1")
	}

	if 1 == 2 {
		fmt.Println("1 is equal to 2")
	} else {
		fmt.Println("1 is not equal to 2")
	}

	const hey = "hey"
	if hey == "Hello" {
		fmt.Println("hello")
	} else if hey == "hi" {
		fmt.Println("hi")
	} else {
		fmt.Println("hey")
	}
}
