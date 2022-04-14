package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.QuoteToASCII(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s)

}
