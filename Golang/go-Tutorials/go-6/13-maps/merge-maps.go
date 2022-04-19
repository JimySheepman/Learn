package main

import "fmt"

func main() {
	first := map[string]int{"a": 1, "b": 2, "c": 3}
	second := map[string]int{"a": 1, "e": 5, "c": 3, "d": 4}

	for k, v := range second {
		first[k] = v
	}

	fmt.Println(first)
}
