package main

import "fmt"

func main() {
	c1 := make(chan int)
	c1 <- 1
	fmt.Println("hello concurrency")
}
