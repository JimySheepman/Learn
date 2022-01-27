package main

import (
	"log"
	"os"
)

func main() {
	oldLocation := "/home/merlins/Projeler/Learn/Golang/go-6/test.txt"
	newLocation := "/home/merlins/Projeler/Learn/Golang/go-6/18-files-and-directories/test.txt"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}
