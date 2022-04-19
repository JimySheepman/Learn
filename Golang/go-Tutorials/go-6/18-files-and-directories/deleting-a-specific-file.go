package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("/home/merlins/Projeler/Learn/Golang/go-6/testing.txt")
	if err != nil {
		log.Fatal(err)
	}
}
