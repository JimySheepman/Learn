package main

import (
	"compress/gzip"
	"log"
	"os"
)

func main() {
	// Create .gz file to write to
	outputFile, err := os.Create("test.txt.gz")
	if err != nil {
		log.Fatal(err)
	}

	// Create a gzip writer on top of file writer
	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	_, err = gzipWriter.Write([]byte("Gophers rule!\n"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Compressed data written to file.")
}
