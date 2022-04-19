package main

import "fmt"

func main() {
	var employee = []string{
		"John",
		"Jane",
		"Jack",
		"Jimmy",
	}
	fmt.Printf("Array: %s\n", employee)

	slice := employee[0:2]
	fmt.Printf("Slice: %s\n", slice)
	slice[0] = "Banana"
	fmt.Printf("Array: %s\n", employee)
	fmt.Printf("Slice: %s\n", slice)

	var slice2 []int
	if slice2 == nil {
		fmt.Println("Slice is nil")
	}
	var slice3 = make([]int, 5, 5)
	fmt.Println(slice3)

	slice4 := employee[2:4]
	subslice := slice4[1:2]
	fmt.Printf("Array: %s\n", employee)
	fmt.Printf("Slice: %s\n", slice4)
	fmt.Printf("Subslice: %s\n", subslice)
	fmt.Printf("Len: %d\nCap: %d\n", len(slice4), cap(slice4))
	slice4 = append(slice4, "Banana-man")
	fmt.Println(slice4)
	fmt.Println(employee)
	result := copy(slice4, employee)
	fmt.Printf("%d values copied\n", result)
}
