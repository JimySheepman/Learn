package main

import "fmt"

func main() {
	johnPrice := computePrice(145.90, 3)
	fmt.Println("John:", johnPrice, "rate:", rate)
}

func computePrice(rate float32, nights int) float32 {
	return rate * float32(nights)
}
