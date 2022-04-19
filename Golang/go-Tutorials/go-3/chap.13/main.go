package main

import (
	"fmt"
	"time"
)

/*
	Predeclared types
		boolean type
			* bool
		String type
			* string
		numeric type
			* uint, uint8,uint32,uint64
			* int, int8,int32,int64
			* float, float32, float64
			* complex64,complex128
	Composite type
		* arrays
		* pointer
		* functions
		* slices
		* map channels
		* struct
		* interface

*/

type test struct {
	Name     string
	Capacity uint8
	Rooms    uint8
	Smoking  bool
}
type Birthdate time.Time
type ExchangeRate map[string]float64

type Hotel struct {
	Name     string
	Capacity uint8
	Rooms    uint8
	Smoking  bool
	Country
}

type Country struct {
	Name        string
	CapitalCity string
}

func main() {
	var rooms uint8 = 130
	var hotelName string = "New Golang Hotel"
	var vacancies bool
	var arr [3]uint8
	var myPointer *uint8
	var roomNumbers []uint8
	var score map[string]uint8
	var received chan<- bool
	fmt.Println(rooms, hotelName, vacancies, arr, myPointer)
	fmt.Println(roomNumbers, score, received)
	france := Country{
		Name:        "France",
		CapitalCity: "Paris",
	}
	fmt.Println(france)

	hotel := Hotel{
		Name:    "Hotel super luxe",
		Country: Country{Name: "France"},
	}
	fmt.Println(hotel.Country.Name)
}
