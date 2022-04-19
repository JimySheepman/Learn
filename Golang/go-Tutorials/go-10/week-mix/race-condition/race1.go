package main

import "fmt"

func main() {
	fmt.Println("---------------------")
	m := make(map[string]string)
	go func() {
		m["1"] = "a"
	}()
	m["2"] = "b"

	for i, v := range m {
		fmt.Println(i, v)
	}
}

// go build -o race1 race1.go && for i in {1..100}; do ./race1; done;
// go build -o race1 race1.go && for i in {1..100}; do ./race1 >> race1.txt; done;
