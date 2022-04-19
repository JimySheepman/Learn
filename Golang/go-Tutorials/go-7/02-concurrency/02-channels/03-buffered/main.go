package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 2)

	ch1 <- "Hello"
	ch1 <- "Hi"

	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

	write := func(ch chan int) {
		for i := 0; i <= 4; i++ {
			ch <- 1
			fmt.Println("Sent:", 1)
		}
		close(ch)
	}

	ch2 := make(chan int, 2)
	go write(ch2)
	time.Sleep(2 * time.Second)
	for v := range ch2 {
		fmt.Println("Received:", v)
		time.Sleep(2 * time.Second)
	}
}
