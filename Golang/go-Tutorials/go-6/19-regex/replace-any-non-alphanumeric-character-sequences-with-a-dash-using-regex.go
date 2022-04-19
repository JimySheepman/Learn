package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	newStr := reg.ReplaceAllString("#Golang#Python$Php&Kotlin@@", "-")
	fmt.Println(newStr)
}
