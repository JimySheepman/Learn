package main

import (
	"fmt"
	c "parent/child"
)

func main() {
	a := &c.Data
	fmt.Println("Address of Data is", a)
	fmt.Println("Value of Data is", *a)
	*a++
	c.PrintData()
}
