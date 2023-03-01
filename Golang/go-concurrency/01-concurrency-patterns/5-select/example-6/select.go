package main

import (
	"fmt"
)

func main() {
	const n = 6
	rightmost := make(chan int)

	var left chan int
	right := rightmost

	for i := 0; i < n; i++ {
		left = make(chan int)
		go pass(left, right)
		right = left
	}

	fmt.Println("Goroutines Are Waiting")

	go func(c chan int) {
		fmt.Println("Give Gopher1 the inital value")
		c <- 1
	}(left)

	fmt.Printf("Final Value: %d\n", <-rightmost)
}

func pass(left, right chan int) {
	value := <-left
	right <- 1 + value

	fmt.Printf("Left[%d] Right[%d]\n", value, value+1)
}
