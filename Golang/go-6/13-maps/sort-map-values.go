package main

import (
	"fmt"
	"sort"
)

func main() {
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}

	values := make([]int, 0, len(unSortedMap))

	for _, v := range unSortedMap {
		values = append(values, v)
	}

	sort.Ints(values)

	for _, v := range values {
		fmt.Println(v)
	}
}
