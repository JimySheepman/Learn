package main

import (
	"fmt"
	"strings"
)

func main() {
	replacer := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(replacer.Replace("This is <b>HTML</b>!"))
}
