package main

import (
	"fmt"
	"sort"
)

func main() {
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}

	keys := make([]string, 0, len(unSortedMap))

	for k := range unSortedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, unSortedMap[k])
	}
}
