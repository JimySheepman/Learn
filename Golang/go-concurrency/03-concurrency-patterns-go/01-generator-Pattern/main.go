package main

import "fmt"

func fib(n int) chan int {
	c := make(chan int)
	go func() {
		for i, j := 0, 1; i < n; i, j = i+j, i {
			c <- i
		}
		close(c)
	}()

	return c
}

func main() {
	for i := range fib(1000) {
		v := i * i
		fmt.Println(v)
	}
}
