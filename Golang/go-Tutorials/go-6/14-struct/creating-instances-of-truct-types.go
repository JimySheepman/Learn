package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
	color   string

	geometry struct {
		area      int
		perimeter int
	}
}

func main() {
	var rect rectangle
	rect.length = 10
	rect.breadth = 20
	rect.color = "Red"

	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * (rect.breadth + rect.length)

	fmt.Println(rect)
	fmt.Println("Area:\t", rect.geometry.area)
	fmt.Println("Perimeter:\t", rect.geometry.perimeter)

}
