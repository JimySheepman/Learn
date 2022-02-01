package main

import "fmt"

func main() {
	hello := "hello"

	switch hello {
	case "hello":
		fmt.Println("hello")
	case "hi":
		fmt.Println("hi")
	default:
		fmt.Println("hey")
	}

	switch hello {
	case "hello":
		fmt.Println("hello")
		fallthrough
	case "hi":
		fmt.Println("hi")
		fallthrough
	default:
		fmt.Println("hey")
	}
}
