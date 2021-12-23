package main

import (
	"fmt"
	"sync"
	"time"
)

func workers(id int) {
	fmt.Printf("Worker %d startinng\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			workers(i)
		}()
	}
	wg.Wait()
}
