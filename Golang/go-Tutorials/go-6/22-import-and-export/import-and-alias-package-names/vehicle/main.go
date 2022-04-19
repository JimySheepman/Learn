package main

import (
	car1 "vehicle/car"
	car2 "vehicle/car"

	"fmt"
)

func main() {
	c1 := new(car1.Car)
	c1.Single(10)
	fmt.Println(c1.Price)

	c2 := new(car2.Car)
	c2.Double(10)
	fmt.Println(c2.Price)
}
