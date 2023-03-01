package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	web1   = fakeSearch("web")
	web2   = fakeSearch("web")
	image1 = fakeSearch("image")
	image2 = fakeSearch("image")
	video1 = fakeSearch("video")
	video2 = fakeSearch("video")
)

type (
	result string
	search func(query string) result
)

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()
	results := google("golang")
	elapsed := time.Since(start)

	fmt.Println(results)
	fmt.Println(elapsed)
}

func fakeSearch(kind string) search {
	return func(query string) result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func google(query string) (results []result) {
	c := make(chan result)

	go func() {
		c <- first(query, web1, web2)
	}()

	go func() {
		c <- first(query, image1, image2)
	}()

	go func() {
		c <- first(query, video1, video2)
	}()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case r := <-c:
			results = append(results, r)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}

	return results
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
