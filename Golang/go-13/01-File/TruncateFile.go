package main

import (
	"log"
	"os"
)

func main() {
	err := os.Truncate("text.txt", 100)
	if err != nil {
		log.Fatal(err)
	}
}
