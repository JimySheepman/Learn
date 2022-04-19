package main

import "fmt"

type rectangle struct {
	length  float64
	breadth float64
	color   string
}

func main() {
	r1 := rectangle{10, 20, "Green"}
	fmt.Println(r1)

	r2 := r1
	r2.color = "Pink"
	fmt.Println(r2)

	r3 := &r1
	r3.color = "Red"
	fmt.Println(r3)

	fmt.Println(r1)
}
