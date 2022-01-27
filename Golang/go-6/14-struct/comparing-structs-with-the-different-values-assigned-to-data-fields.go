package main

import "fmt"

type rectangle struct {
	length  float64
	breadth float64
	color   string
}

func main() {
	var rect1 = rectangle{10, 20, "Green"}
	rect2 := rectangle{length: 20, breadth: 10, color: "Red"}

	if rect1 == rect2 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	rect3 := new(rectangle)
	var rect4 = &rectangle{}

	if rect3 == rect4 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
