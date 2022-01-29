package main

import "fmt"

type Triangle struct {
	base, height float32
}
type Square struct {
	length float32
}

type Rectangle struct {
	length, width float32
}

type Circle struct {
	radius float32
}

func (t Triangle) Area() float32 {
	return 0.5 * t.base * t.height
}

func (s Square) Area() float32 {
	return s.length * s.length
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func (c Circle) Area() float32 {
	return 3.14 * (c.radius * c.radius)
}

type Area interface {
	Area() float32
}

func main() {
	t := Triangle{base: 15, height: 25}
	s := Square{length: 5}
	r := Rectangle{length: 5, width: 10}
	c := Circle{radius: 5}

	// define a interface
	var a Area

	a = t
	fmt.Println("Area of Triangle", a.Area())

	a = s
	fmt.Println("Area of Square", a.Area())

	a = r
	fmt.Println("Area of Rectangle", a.Area())

	a = c
	fmt.Println("Area of Circle", a.Area())
}
