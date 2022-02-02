package main

import "fmt"

func main() {
	items := []int{9, 5, 7, 3}

	sorted := QuickSort(items)

	fmt.Println("sorted:", sorted)
}

func QuickSort(items []int) []int {
	low := 0
	high := len(items) - 1

	quickSort(items, low, high)

	return items
}

func quickSort(items []int, low, high int) {
	if low < high {
		pivot := Partition(items, low, high)

		quickSort(items, low, pivot-1)
		quickSort(items, pivot+1, high)
	}
}

func Partition(items []int, low, high int) int {
	pivot := items[high]
	i := low

	for j := low; j < high; j++ {
		if items[j] < pivot {
			items[i], items[j] = items[j], items[i]
			i++
		}
	}
	items[i], items[high] = items[high], items[i]

	return i
}
