package main

import "fmt"

func main() {
	unbufferedChan := make(chan int)

	go func(unbufChan chan int) {
		value := <-unbufChan
		fmt.Println(value)
	}(unbufferedChan)

	go func(unbufChan chan int) {
		unbufChan <- 1
	}(unbufferedChan)

	fmt.Println("hello channels")

}
