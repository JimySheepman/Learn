package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40}
	a = append(a, 50)
	fmt.Println(a)

	s := []uint{10, 20, 30, 40}
	fmt.Printf("Length : %d - Capacity : %d\n", len(s), cap(s))
	s = append(s, 50)
	fmt.Printf("Length : %d - Capacity : %d\n", len(s), cap(s))
	s = append(s, 60)
	fmt.Printf("Length : %d - Capacity : %d\n", len(s), cap(s))
	s = append(s, 70)
	fmt.Printf("Length : %d - Capacity : %d\n", len(s), cap(s))
	s = append(s, 80)
	fmt.Printf("Length : %d - Capacity : %d\n", len(s), cap(s))
	s = append(s, 90)
	fmt.Printf("Length : %d - Capacity : %d\n", len(s), cap(s))
}
