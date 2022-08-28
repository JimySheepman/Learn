package main

import "fmt"

func main() {
	const hotelName string = "Gopher Hotel"
	const longitude = 24.806078
	const latitude = -78.243027
	var occupancy int = 12
	fmt.Println(hotelName, longitude, latitude)
	fmt.Println(occupancy)
}
