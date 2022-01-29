package main

import "fmt"

var (
	a = 654
	b = false
	c = 2.651
	d = 4 + 1i
)

func main() {
	fmt.Printf("d for Integer: %d\n", a)
	fmt.Printf("t for Boolean: %t\n", b)
	fmt.Printf("g for Float: %g\n", c)
	fmt.Printf("e for scientific Notation: %e\n", d)
}
