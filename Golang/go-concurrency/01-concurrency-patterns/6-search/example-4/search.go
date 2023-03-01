package main

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	result string
	search func(query string) result
)

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()

	result := first("golang",
		fakeSearch("replica 1"),
		fakeSearch("replica 2"))

	elapsed := time.Since(start)

	fmt.Println(result)
	fmt.Println(elapsed)
}

func fakeSearch(kind string) search {
	return func(query string) result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func first(query string, replicas ...search) result {
	c := make(chan result)

	searchReplica := func(i int) {
		c <- replicas[i](query)
	}

	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}
