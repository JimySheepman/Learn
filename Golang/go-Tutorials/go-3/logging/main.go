package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	log.Println("System is starting")
	log.Println("Retrieving user 12")
	log.Println("Compute the total of invoice 1262663663")
}
