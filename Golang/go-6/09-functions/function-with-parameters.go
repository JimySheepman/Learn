package main

import "fmt"

func add(x, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

func main() {
	add(20, 30)
}
