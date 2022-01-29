package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "nEw strings"
	str = strings.ToLower(str)
	result := strings.ReplaceAll(str, " ", "-")
	fmt.Println(result + ".go")
}
