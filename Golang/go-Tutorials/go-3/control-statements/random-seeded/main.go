package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	hotelName := "The Gopher Hotel"
	fmt.Println("Hotel " + hotelName)
	var roomsAvailable int
	rand.Seed(time.Now().UTC().UnixNano())
	var rooms, roomsOccupied int = 100, rand.Intn(100)
	roomsAvailable = rooms - roomsOccupied
	fmt.Println(roomsAvailable, "rooms available")
}
