package main

import (
	h "analysis/human"
	v "analysis/vehicle"

	"fmt"
)

func main() {
	var bmw v.Vehicle
	bmw = v.Car("World Top Brand")

	var labour h.Human
	labour = h.Man("Software Developer")

	for i, j := range bmw.Structure() {
		fmt.Printf("%-15s <=====> %15s\n", j, labour.Structure()[i])
	}
}
