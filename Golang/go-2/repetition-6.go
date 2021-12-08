package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}
type Circle struct {
	x, y, radius float64
}
type Rectangle struct {
	width, height float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}
func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}
func getArea(shape Shape) float64 {
	return shape.area()
}
func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}
func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}
	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))

	result, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	result, err = Sqrt(9)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
