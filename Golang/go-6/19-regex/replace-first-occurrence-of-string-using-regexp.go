package main

import (
	"fmt"
	"regexp"
)

func main() {
	strEx := "Php-Golang-Php-Python-Php-Kotlin"
	reStr := regexp.MustCompile("^(.*?)Php(.*)$")
	repStr := "${1}Java$2"
	output := reStr.ReplaceAllString(strEx, repStr)
	fmt.Println(output)
}
