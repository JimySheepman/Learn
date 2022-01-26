package main

import "fmt"

func main() {

	variadicExample()
	variadicExample("red", "blue")
	variadicExample("red", "blue", "green")
	variadicExample("red", "blue", "green", "yellow")
}

func variadicExample(s ...string) {
	fmt.Println(s)
}
