package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "Hello X42 I'm a Y-32.35 string Z30"

	re := regexp.MustCompile(`[A-Z][^A-Z]*`)

	fmt.Printf("Pattern: %v\n", re.String()) // Print Pattern

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
