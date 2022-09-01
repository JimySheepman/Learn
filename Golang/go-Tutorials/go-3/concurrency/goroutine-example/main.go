package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("launch first goroutine")
	go printNumber()
	fmt.Println("launch second goroutine")
	go printNumber()
	time.Sleep(1 * time.Minute)
}

func printNumber() {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		i++
		fmt.Println(i)
	}
}
