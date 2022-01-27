package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "We @@@Love@@@@ #Go!$! ****Programming****Language^^^"

	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	fmt.Printf("Pattern: %v\n", re.String())
	fmt.Println(re.MatchString(str1))

	submatchall := re.FindAllString(str1, -1)

	for _, v := range submatchall {
		fmt.Println(v)
	}
}
