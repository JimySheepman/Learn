package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.csv")
	defer file.Close()
	if err != nil {
		log.Printf("impossible to open file %s", err)
		return
	}

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}
