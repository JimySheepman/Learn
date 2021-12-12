package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var rooms, roomsOccupied int = 100, rand.Intn(100)
	fmt.Println("rooms :", rooms, " roomsOccupied :", roomsOccupied)
	fmt.Println("boolean expression : 'rooms > roomsOccupied' is equal to :")
	fmt.Println(rooms > roomsOccupied)
	fmt.Println("boolean expression : 'roomsOccupied > rooms' is equal to :")
	fmt.Println(roomsOccupied > rooms)
	fmt.Println("boolean expression : 'roomsOccupied == rooms' is equal to :")
	fmt.Println(roomsOccupied == rooms)
	a := 21
	b := 42

	fmt.Println(a == b)
	fmt.Println(a == a)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
	fmt.Println(a <= a)
	fmt.Println(a >= a)
	var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)

	fmt.Println("John is", ageJohn, "years old")
	fmt.Println("Paul is", agePaul, "years old")

	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")
	} else {
		if agePaul == ageJohn {
			fmt.Println("Paul and John have the same age")
		} else {
			fmt.Println("Paul is younger than John")
		}
	}
	switch ageJohn {
	case 10:
		fmt.Println("John is 10 years old")
	case 20:
		fmt.Println("John is 20 years old")
	case 100:
		fmt.Println("John is 100 years old")
	default:
		fmt.Println("John has neither 10,20 nor 100 years old")
	}

	switch ageSum := ageJohn + agePaul; ageSum {
	case 10:
		fmt.Println("ageJohn + agePaul = 10")
	case 20, 30, 40:
		fmt.Println("ageJohn + agePaul = 20 or 30 or 40")
	case 2 * agePaul:
		fmt.Println("ageJohn + agePaul = 2 times agePaul")
	}

	switch {
	case agePaul > ageJohn:
		fmt.Println("agePaul > ageJohn")
	case agePaul == ageJohn:
		fmt.Println("agePaul == ageJohn")
	case agePaul < ageJohn:
		fmt.Println("agePaul < ageJohn")
	}
	const emailToSend = 3
	emailSent := 0

	for emailSent < emailToSend {
		fmt.Println("sending email..")
		emailSent++ //*\label{forWithSingle2}
	}
	fmt.Println("end of program")
	ageJohn = rand.Intn(110)
	fmt.Println("John is", ageJohn, "years old")

	for i := 0; i < ageJohn; i++ {
		fmt.Println("interation N", i)
	}
}
