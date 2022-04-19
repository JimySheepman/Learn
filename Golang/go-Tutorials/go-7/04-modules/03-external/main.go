package main

import (
	clients "external-sample/pkg"
	"fmt"
)

func main() {
	a := 3
	b := 5
	sum := clients.Sum(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, sum)
}
