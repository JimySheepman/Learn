package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	str1 := "how much for the maple syrup? $20.99? That's ridiculous!!!"
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	str1 = re.ReplaceAllString(str1, " ")
	fmt.Println(str1)
}
