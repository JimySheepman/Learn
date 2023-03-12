package main

import (
	"strconv"
	"sync"
)

const (
	cycles        = 15
	calssPerCycle = 100
)

func RunMockServer(cache *KeyStoreCache) {
	var wg sync.WaitGroup

	for c := 0; c < cycles; c++ {
		wg.Add(1)
		go func() {
			for i := 0; i < calssPerCycle; i++ {
				cache.Get("Test" + strconv.Itoa(i))
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
