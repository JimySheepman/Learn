package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		fmt.Println("hello from go routine")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("hello concurrency")
}
