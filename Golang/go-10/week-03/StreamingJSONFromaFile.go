package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	fileReader, _ := os.Open("week-03/output.json")
	var tpphu map[string]interface{}
	json.NewDecoder(fileReader).Decode(&tpphu)
	fmt.Println(tpphu)

}
