package main

import "fmt"

func main() {
	intArray := [5]int{10, 20, 30, 40, 50}

	fmt.Println("\nExample 1")
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	fmt.Println("\nExample 2")
	for index, element := range intArray {
		fmt.Println(index, "=>", element)

	}

	fmt.Println("\nExample 3")
	for _, value := range intArray {
		fmt.Println(value)
	}

	j := 0
	fmt.Println("\nExample 4")
	for range intArray {
		fmt.Println(intArray[j])
		j++
	}
}
