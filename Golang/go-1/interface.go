package main

import "fmt"

func main() {

	// struct objects
	var s1 Sparrow
	var p1 Plane

	// call methods
	s1.fly()
	p1.fly()
}

type Flyable interface {
	fly()
}

type Sparrow struct{}
type Plane struct{}

func (s Sparrow) fly() {

	fmt.Println("Sparrow flies at a speed of 10")
}

func (p Plane) fly() {

	fmt.Println("Plane flies at a speed of 2000")
}
