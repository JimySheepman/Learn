package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "Split   String on \nwhite    \tspaces."

	re := regexp.MustCompile(`\S+`)

	fmt.Printf("Pattern: %v\n", re.String()) // Print Pattern

	fmt.Printf("String contains any match: %v\n", re.MatchString(str1)) // True

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
