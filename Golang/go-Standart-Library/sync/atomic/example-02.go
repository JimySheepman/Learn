package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	type Map map[string]string
	var m atomic.Value
	m.Store(make(Map))
	var mu sync.Mutex
	read := func(key string) (val string) {
		m1 := m.Load().(Map)
		return m1[key]
	}

	insert := func(key, val string) {
		mu.Lock()
		defer mu.Unlock()
		m1 := m.Load().(Map)
		m2 := make(Map)
		for k, v := range m1 {
			m2[k] = v
		}
		m2[key] = val
		m.Store(m2)
	}
	_, _ = read, insert
}
