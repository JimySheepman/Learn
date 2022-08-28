package main

import "fmt"

func main() {
	const occupancyLimit = 12

	var occupancyLimit1 uint8
	var occupancyLimit2 int64
	var occupancyLimit3 float32

	occupancyLimit1 = occupancyLimit
	occupancyLimit2 = occupancyLimit
	occupancyLimit3 = occupancyLimit

	fmt.Println(occupancyLimit1, occupancyLimit2, occupancyLimit3)
}
