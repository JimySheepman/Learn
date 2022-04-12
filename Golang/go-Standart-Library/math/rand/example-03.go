package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(86)
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
}
