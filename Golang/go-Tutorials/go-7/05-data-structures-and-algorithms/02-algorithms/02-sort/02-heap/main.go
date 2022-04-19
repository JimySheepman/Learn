package main

import "fmt"

func main() {
	items := []int{9, 5, 7, 3}

	sorted := HeapSort(items)

	fmt.Println("sorted:", sorted)
}

func HeapSort(items []int) []int {
	length := len(items)

	for i := length / 2; i >= 0; i-- {
		Heap(items, i, length)
	}

	for i := length; i > 1; i-- {
		items[0], items[i-1] = items[i-1], items[0]
		Heap(items, 0, i-1)
	}
	return items
}

func Heap(items []int, root, length int) {
	max := root
	left := root*2 + 1
	right := root*2 + 2

	if left < length && items[left] > items[max] {
		max = left
	}
	if right < length && items[right] > items[max] {
		max = right
	}
	if max != root {
		items[root], items[max] = items[max], items[root]
		Heap(items, max, length)
	}
}
