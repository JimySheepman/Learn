package main

import (
	"log"
)

func main() {
	ch := make(chan int, 1)
	go dummy(ch)
	log.Println("waiting for reception...")
	ch <- 45
	ch <- 58
	ch <- 100
}

func dummy(c chan int) {
	smth := <-c
	log.Println("has received something", smth)
}
