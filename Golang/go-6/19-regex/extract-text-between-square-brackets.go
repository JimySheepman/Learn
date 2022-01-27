package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str1 := "this is a [sapmle] [[string]] with [SOME] special words"
	re := regexp.MustCompile(`\[([^\[\]]*)\]`)
	fmt.Printf("Pattern: %v\n", re.String())
	fmt.Println("Matched:", re.MatchString(str1))

	fmt.Println("\nText between square brackets:")
	submatchall := re.FindAllString(str1, -1)
	for _, v := range submatchall {
		v = strings.Trim(v, "[")
		v = strings.Trim(v, "]")
		fmt.Println(v)
	}
}
