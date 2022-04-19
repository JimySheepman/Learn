package main

import "fmt"

func main() {
	a := []int{5, 6, 7} // Create a smaller slice
	fmt.Printf("[Slice:A] Length is %d Capacity is %d\n", len(a), cap(a))

	b := make([]int, 5, 10) // Create a bigger slice
	copy(b, a)              // Copy function
	fmt.Printf("[Slice:B] Length is %d Capacity is %d\n", len(b), cap(b))

	fmt.Println("Slice B after copying:", b)
	b[3] = 8
	b[4] = 9
	fmt.Println("Slice B after adding elements:", b)
}
