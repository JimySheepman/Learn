package main

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
)

func listFiles(file *zip.File) error {
	fileread, err := file.Open()
	if err != nil {
		msg := "Failed to open zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}
	defer fileread.Close()

	fmt.Fprintf(os.Stdout, "%s:", file.Name)

	if err != nil {
		msg := "Failed to read zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}

	fmt.Println()

	return nil
}

func main() {
	read, err := zip.OpenReader("test.zip")
	if err != nil {
		msg := "Failed to open: %s"
		log.Fatalf(msg, err)
	}
	defer read.Close()

	for _, file := range read.File {
		if err := listFiles(file); err != nil {
			log.Fatalf("Failed to read %s from zip: %s", file.Name, err)
		}
	}
}
