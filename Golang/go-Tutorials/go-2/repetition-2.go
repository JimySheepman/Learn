package main

import "fmt"

func main() {
	var a, b int = 5, 3

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	fmt.Println(a == b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!(true && false))
	A, B := 60, 13
	fmt.Println(A & B)
	fmt.Println(A | B)
	fmt.Println(A ^ B)
	fmt.Println(^A)
	fmt.Println(A << 2)
	fmt.Println(A >> 2)

}
