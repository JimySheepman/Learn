package main

import "fmt"

func main() {

	strArray1 := [3]string{"Japan", "Australia", "Germany"}
	strArray2 := strArray1
	strArray3 := &strArray1

	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)

	strArray1[0] = "Canada"

	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)
	fmt.Printf("strArray3: %v\n", *strArray3)
}
