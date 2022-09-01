package main

import (
	"time"
)

func main() {
	syncCh := make(chan bool)
	go func() {
		longTask2()
		syncCh <- true
	}()
	longTask()
	<-syncCh
}

func longTask2() {
	time.Sleep(1 * time.Second)
}

func longTask() {
	time.Sleep(3 * time.Second)
}
