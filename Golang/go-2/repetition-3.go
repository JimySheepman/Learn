package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200
	var ret int
	ret = max(a, b)
	fmt.Printf("Max value is : %d\n", ret)

	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
	var A int = 20 /* actual variable declaration */
	var ip *int    /* pointer variable declaration */
	ip = &A        /* store address of a in pointer variable*/
	fmt.Printf("Address of a variable: %x\n", &A)
	fmt.Printf("Address stored in ip variable: %x\n", ip)
	fmt.Printf("Value of *ip variable: %d\n", *ip)
}

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
