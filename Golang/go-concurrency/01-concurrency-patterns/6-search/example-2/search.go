package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	web   = fakeSearch("web")
	image = fakeSearch("image")
	video = fakeSearch("video")
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
		c <- web(query)
	}()

	go func() {
		c <- image(query)
	}()

	go func() {
		c <- video(query)
	}()

	for i := 0; i < 3; i++ {
		r := <-c
		results = append(results, r)
	}

	return results
}
