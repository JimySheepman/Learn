package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsRune("aardvark", 97))
	fmt.Println(strings.ContainsRune("timeout", 97))
}
