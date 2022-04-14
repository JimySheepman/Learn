package main

import (
	"fmt"
	"regexp"
)

func main() {
	content := []byte(`
	# comment line
	option1: value1
	option2: value2
`)

	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	loc := pattern.FindIndex(content)
	fmt.Println(loc)
	fmt.Println(string(content[loc[0]:loc[1]]))
}