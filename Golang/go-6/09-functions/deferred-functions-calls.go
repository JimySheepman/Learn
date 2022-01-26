package main

import "fmt"

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

func main() {
	defer second()
	first()
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
