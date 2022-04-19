package main

import "fmt"

func main() {
	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}

	fmt.Println("\n---------------Example 1 --------------------")
	for index, element := range strSlice {
		fmt.Println(index, "--", element)
	}

	fmt.Println("\n---------------Example 2 --------------------")
	for _, value := range strSlice {
		fmt.Println(value)
	}

	j := 0
	fmt.Println("\n---------------Example 3 --------------------")
	for range strSlice {
		fmt.Println(strSlice[j])
		j++
	}

	fmt.Println("\n---------------Example 4 --------------------")
	var slice1 = []string{"india", "japan", "canada"}
	var slice2 = []string{"australia", "russia"}

	slice2 = append(slice2, slice1...)
	fmt.Println(slice2)
}
