package main

import (
	"fmt"
	"sort"
)

func main() {
	scores := []int{10, 89, 76, 3, 20, 12}

	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })
	fmt.Println(scores)

	sort.Slice(scores, func(i, j int) bool { return scores[i] > scores[j] })
	fmt.Println(scores)
}
