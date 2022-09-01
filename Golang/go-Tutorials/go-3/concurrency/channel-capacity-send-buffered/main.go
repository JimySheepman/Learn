package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go dummy(ch)
	log.Println("waiting for reception...")
	ch <- 45
	log.Println("received")
}

func dummy(c chan int) {
	time.Sleep(3 * time.Second)
	<-c
}
