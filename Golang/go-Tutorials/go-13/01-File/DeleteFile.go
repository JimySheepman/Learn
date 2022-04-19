package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("test2.txt")
	if err != nil {
		log.Fatal(err)
	}
}
