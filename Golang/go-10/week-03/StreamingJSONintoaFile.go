package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
}

func main() {
	tpphu := Person{
		Name:   "Tran Phong Phu",
		Age:    20,
		Emails: []string{"vietean@gmail.com", "phu.tran@nordiccoder.com"},
	}

	fileWriter, _ := os.Create("output.json")
	json.NewEncoder(fileWriter).Encode(tpphu)
}
