package main

import "fmt"

func ReverseArray(array []int32) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

func main() {
	s := []int32{4, 4, 4, 3, 4, 4, 4, 4, 2, 2, 2, 1, 1, 5, 5}
	ReverseArray(s)
	fmt.Println(s)
}
