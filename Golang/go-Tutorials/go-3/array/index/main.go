package main

import "fmt"

func main() {
	myArray := [2]int{156, 147}

	for index, element := range myArray {
		fmt.Printf("element at index %d is %d\n", index, element)
	}

	for _, element := range myArray {
		fmt.Println(element)
	}

	for index, _ := range myArray {
		fmt.Println(index)
	}

	myArray2 := [2]int{156, 147}
	for i := 0; i < len(myArray2); i++ {
		fmt.Printf("element at index %d is %d\n", i, myArray2[i])
	}
}

func getIndex(haystack [10]int, needle int) int {
	for index, element := range haystack {
		if element == needle {
			return index
		}
	}
	return -1
}
