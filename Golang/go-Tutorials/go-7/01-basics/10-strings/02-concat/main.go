package main

import (
	"fmt"
	"strconv"
)

func main() {
	firstName, lastName := "Go", "Gopher"
	fullName := firstName + lastName
	fmt.Printf("Full Name: %s\n", fullName)

	fullName += " (o_0)"
	fmt.Printf("Full Name: %s\n", fullName)

	str := strconv.FormatInt(1, 2) + " is " + strconv.FormatBool(true)
	fmt.Println(str)
}
