package main

import "fmt"

func main() {
	printRoomDetails2(112, 3, 4)
}

func printRoomDetails2(roomNumber, size, nights int) {
	if nights == 1 {
		fmt.Println(roomNumber, ":", size, "people /", nights, "night")
	} else {
		fmt.Println(roomNumber, ":", size, "people /", nights, "nights")
	}
}
