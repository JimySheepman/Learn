package main

import (
	"fmt"
	"sync"
)

func main() {
	var i = 0

	lock := sync.Mutex{}

	lock.Lock()
	lock.Lock()
	i++
	lock.Unlock()
	fmt.Println("hello concurrency")
}