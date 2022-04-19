package main

import (
	"log"
	"os"
)

func main() {
	oldName := "empty.txt"
	newName := "testing.txt"
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}
