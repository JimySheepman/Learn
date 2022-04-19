package main

import (
	"fmt"
	"time"
)

func main() {
	go printLoop()

	fmt.Println("Don't wait for it")

	go func() {
		fmt.Println("new func")
	}()

	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("Run: %d\n", i)
		}(i)
	}
	time.Sleep(2 * time.Second)
}

func printLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}
