package main

import (
	"fmt"
	"time"
)

func main() {
	unbufferedChan := make(chan int)

	go func(unbufChan chan int) {
		value := <-unbufChan
		fmt.Println(value)
	}(unbufferedChan)

	go func(unbufChan chan int) {
		unbufChan <- 1
	}(unbufferedChan)

	time.Sleep(time.Second)

	fmt.Println("hello channels")

}
