package main

import "fmt"

func main() {
	items := []int{9, 5, 7, 3}

	sorted := InsertionSort(items)

	fmt.Println("sorted:", sorted)
}

func InsertionSort(items []int) []int {
	length := len(items)

	for i := 0; i < length; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j--
		}
	}
	return items
}
