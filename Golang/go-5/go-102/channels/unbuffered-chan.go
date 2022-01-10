package main

import "fmt"

func main() {
	unbufferedChan := make(chan int)

	var unbufferedChan2 chan int

	fmt.Println(unbufferedChan)
	fmt.Println(unbufferedChan2)

	go func(unbufChan <-chan int) {
		value := <-unbufChan
		fmt.Println(value)

	}(unbufferedChan)

	unbufferedChan <- 1
}
