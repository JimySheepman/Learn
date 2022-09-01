package main

import (
	"errors"
	"log"
)

func main() {
	err := transferFileContents("/my/imaginary/file")
	var readingError *ReadingError
	if errors.As(err, &readingError) {
		log.Fatalf("error of reading occured: %s: %s", readingError, readingError.Unwrap())
	}
	var writingError *WritingError
	if errors.As(err, &writingError) {
		log.Fatalf("error of writing occured: %s", err)
	}
	log.Println("transfer done")
}
