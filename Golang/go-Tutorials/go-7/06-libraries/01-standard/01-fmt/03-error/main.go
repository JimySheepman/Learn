package main

import (
	"fmt"
)

func main() {
	Sample("Errorf", Errorf)
}

func Errorf() {
	// creates an error message wit format
	err := fmt.Errorf("user is no found with id: %d", 100)
	fmt.Println(err)
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
