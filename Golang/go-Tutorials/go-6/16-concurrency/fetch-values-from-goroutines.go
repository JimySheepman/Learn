package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func responseSize(url string, nums chan int) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	nums <- len(body)
}

func main() {
	nums := make(chan int)
	wg.Add(1)
	go responseSize("https://www.golangprograms.com", nums)
	fmt.Println(<-nums)
	wg.Wait()
	close(nums)
}
