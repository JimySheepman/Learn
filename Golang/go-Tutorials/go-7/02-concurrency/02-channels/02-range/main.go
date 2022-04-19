package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 100)

	go calc(ch)

	for i := range ch {
		fmt.Printf("%d", i)
	}
}

func calc(ch chan int) {
	for i := 0; i <= 10; i++ {
		ch <- 7
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}
