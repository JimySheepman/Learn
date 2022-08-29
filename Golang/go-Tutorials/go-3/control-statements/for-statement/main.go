package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const emailToSend = 3
	emailSent := 0

	for emailSent < emailToSend {
		fmt.Println("sending email..")
		emailSent++
	}
	fmt.Println("end of program")

	rand.Seed(time.Now().UTC().UnixNano())
	var ageJohn int = rand.Intn(110)
	fmt.Println("John is", ageJohn, "years old")

	for i := 0; i < ageJohn; i++ {
		fmt.Println("interation N", i)
	}
}
