package main

import "fmt"

func main() {
	items := []int{9, 5, 7, 3}

	sorted := SelectionSort(items)

	fmt.Println("sorted:", sorted)
}

func SelectionSort(items []int) []int {
	length := len(items)

	for i := 0; i < length-1; i++ {
		min := 1
		for j := i + 1; j < length; j++ {
			if items[j] < items[min] {
				min = j
			}
		}
		items[i], items[min] = items[min], items[i]
	}
	return items
}
