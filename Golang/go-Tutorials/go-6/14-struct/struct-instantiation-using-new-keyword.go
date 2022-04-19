package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
	color   string
}

func main() {
	rect1 := new(rectangle) // rect1 is a pointer to an instance of rectangle
	rect1.length = 10
	rect1.breadth = 20
	rect1.color = "Green"
	fmt.Println(rect1)

	var rect2 = new(rectangle) // rect2 is an instance of rectangle
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2)
}
