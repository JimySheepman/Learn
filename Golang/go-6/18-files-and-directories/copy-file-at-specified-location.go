package main

import (
	"io"
	"log"
	"os"
)

func main() {
	sourceFile, err := os.Open("/home/merlins/Projeler/Learn/Golang/go-6/testing.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	newFile, err := os.Create("/home/merlins/Projeler/Learn/Golang/go-6/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesCopied)

}
