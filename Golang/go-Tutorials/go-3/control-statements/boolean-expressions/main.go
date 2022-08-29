package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var rooms, roomsOccupied int = 100, rand.Intn(100)

	fmt.Println("rooms : ", rooms, " roomsOccupied : ", roomsOccupied)

	fmt.Println("boolean expression : 'rooms > roomsOccupied' is equal to :", rooms > roomsOccupied)
	fmt.Println("boolean expression : 'roomsOccupied > rooms' is equal to :", rooms < roomsOccupied)
	fmt.Println("boolean expression : 'roomsOccupied == rooms' is equal to :", rooms == roomsOccupied)

}
