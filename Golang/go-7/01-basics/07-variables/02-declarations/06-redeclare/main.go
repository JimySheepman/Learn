package main

import "fmt"

func main() {
	var year int = 2000
	year = 2000 + 200
	var month = 5
	month, day := 5, 20

	fmt.Println(day, month, year)

	var first, second = 1, 2

	first, second = second, first
	fmt.Println(first, second)

}
