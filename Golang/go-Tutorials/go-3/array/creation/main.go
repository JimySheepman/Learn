package main

import "fmt"

func main() {
	var myArray [2]int
	myArray[0] = 156
	myArray[1] = 147

	fmt.Println(myArray)

	myArray2 := [2]int{156, 147}

	fmt.Println(myArray2)

	myArray3 := [...]int{156, 147}

	fmt.Println(myArray3)
}
