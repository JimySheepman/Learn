package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	paper = iota
	grass
	match
)

var smokeMap = map[int]string{
	paper: "paper",
	grass: "grass",
	match: "match",
}

var names = map[int]string{
	paper: "Sandy",
	grass: "Apple",
	match: "Daisy",
}

type Table struct {
	paper chan int
	grass chan int
	match chan int
}

func arbitrate(t *Table, smokers [3]chan int) {
	for {
		time.Sleep(time.Millisecond * 500)
		next := rand.Intn(3)
		fmt.Printf("Table chooses %s: %s\n", smokeMap[next], names[next])
		switch next {
		case paper:
			t.grass <- 1
			t.match <- 1
		case grass:
			t.paper <- 1
			t.match <- 1
		case match:
			t.grass <- 1
			t.paper <- 1
		}
		for _, smoker := range smokers {
			smoker <- next
		}
		wg.Add(1)
		wg.Wait()
	}
}

func smoker(t *Table, name string, smokes int, signal chan int) {
	var chosen = -1
	for {
		chosen = <-signal // blocks

		if smokes != chosen {
			continue
		}

		fmt.Printf("Table: %d grass: %d match: %d\n", len(t.paper), len(t.grass), len(t.match))
		select {
		case <-t.paper:
		case <-t.grass:
		case <-t.match:
		}
		fmt.Printf("Table: %d grass: %d match: %d\n", len(t.paper), len(t.grass), len(t.match))
		time.Sleep(10 * time.Millisecond)
		select {
		case <-t.paper:
		case <-t.grass:
		case <-t.match:
		}
		fmt.Printf("Table: %d grass: %d match: %d\n", len(t.paper), len(t.grass), len(t.match))
		fmt.Printf("%s smokes a cigarette\n", name)
		time.Sleep(time.Millisecond * 500)
		wg.Done()
		time.Sleep(time.Millisecond * 100)
	}
}

const LIMIT = 1

var wg *sync.WaitGroup

func main() {
	wg = new(sync.WaitGroup)
	table := new(Table)
	table.match = make(chan int, LIMIT)
	table.paper = make(chan int, LIMIT)
	table.grass = make(chan int, LIMIT)
	var signals [3]chan int
	// three smokers
	for i := 0; i < 3; i++ {
		signal := make(chan int, 1)
		signals[i] = signal
		go smoker(table, names[i], i, signal)
	}
	fmt.Printf("%s, %s, %s, sit with \n%s, %s, %s\n\n", names[0], names[1], names[2], smokeMap[0], smokeMap[1], smokeMap[2])
	arbitrate(table, signals)
}
