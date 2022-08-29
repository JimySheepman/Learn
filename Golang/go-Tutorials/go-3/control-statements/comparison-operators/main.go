package main

import "fmt"

func main() {
	a := 21
	b := 42

	fmt.Println(a == b)
	fmt.Println(a == a)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
	fmt.Println(a <= a)
	fmt.Println(a >= a)
}
