package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	var w8 sync.WaitGroup
	w8.Add(10)
	for k := 0; k < 10; k++ {
		go caller(&w8)
	}
	w8.Wait()
}

func caller(w8 *sync.WaitGroup) {
	for k := 0; k < 100; k++ {
		res, err := http.Get("http://localhost:8090/status")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		s, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(s))
	}
	w8.Done()
}
