package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)

	fmt.Println("John is", ageJohn, "years old")
	fmt.Println("Paul is", agePaul, "years old")
	fmt.Println("It is", ageJohn > agePaul, "that John is older than Paul")
	fmt.Println("It is", ageJohn == agePaul, "that John and Paul have the same age")
}
