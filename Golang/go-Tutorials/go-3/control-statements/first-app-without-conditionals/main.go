package main

import "fmt"

func main() {
	hotelName := "The Gopher Hotel"
	fmt.Println("Hotel " + hotelName)

	var roomsAvailabe uint8
	var rooms, roomsOccupied uint8 = 100, 10
	roomsAvailabe = rooms - roomsOccupied
	fmt.Println(roomsAvailabe, " rooms available")
}
