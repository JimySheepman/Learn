package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 3.0, 4.0
	h := math.Hypot(a, b)

	fmt.Print("The vector(", a, b, ") has length", h, ".\n")
	fmt.Println("The vector (", a, b, ") has length", h, ".")
	fmt.Printf("The vector (%g %g) has length %g.\n", a, b, h)
}
