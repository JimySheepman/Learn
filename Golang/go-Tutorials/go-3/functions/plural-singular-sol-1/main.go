package main

import "fmt"

func main() {
	printRoomDetails(112, 2, 2)
}

func printRoomDetails(roomNumber, size, nights int) {
	nightText := "nights"
	if nights == 1 {
		nightText = "night"
	}
	fmt.Println(roomNumber, ":", size, "people /", nights, nightText)
}
