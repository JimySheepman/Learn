package main

import "fmt"

func main() {
	var roomNumber, floorNumber int = 154, 3
	fmt.Println(roomNumber, floorNumber)

	var password = "notSecured"
	fmt.Println(password)

	const occupancyLimit = 12
	var occupancyLimit1 uint8
	var occupancyLimit2 int64
	var occupancyLimit3 float32
	occupancyLimit1 = occupancyLimit
	occupancyLimit2 = occupancyLimit
	occupancyLimit3 = occupancyLimit
	fmt.Println(occupancyLimit1, occupancyLimit2, occupancyLimit3)
	const hotelName string = "Gopher Hotel"
	const longitude = 24.806078
	const latitude = -78.243027
	var occupancy int = 12
	fmt.Println(hotelName, longitude, latitude)
	fmt.Println(occupancy)
}
