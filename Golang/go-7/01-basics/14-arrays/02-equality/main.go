package main

import "fmt"

func main() {
	first := [...]int{1, 2}
	second := [...]int{1, 2}
	fmt.Printf("%v==%v => %v\n", first, second, first == second)

	fourth := [2]int{2, 1}
	fmt.Printf("%v==%v => %v\n", first, fourth, first == fourth)
}
