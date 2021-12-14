package main

import (
	"fmt"
	"main/invoice"
)

func init() {
	fmt.Println("main")
}
func main() {
	fmt.Println("--program start--")
	invoice.Print()
}
