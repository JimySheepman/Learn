package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int32
	wg      sync.WaitGroup
)

func main() {
	wg.Add(3)

	go increment("Python")
	go increment("Java")
	go increment("Golang")

	wg.Wait()
	fmt.Println("Counter:", counter)
}

func increment(name string) {
	defer wg.Done()

	for range name {
		atomic.AddInt32(&counter, 1)
		runtime.Gosched()
	}
}
