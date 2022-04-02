package main

import (
	"fmt"
)

var x = 42

func foo() {
	fmt.Println(x)
}

func foo2() {
	fmt.Println(x)
}

func increment() int {
	x++
	return x
}

func package_scope() {
	fmt.Println(x)
	foo()
	fmt.Println(MyName)
	PrintVar()
}

func block_scope() {
	x := 42
	fmt.Println(x)
	foo2()
	{
		fmt.Println(x)
		y := "the credit velongs with the one who is in the ring."
		fmt.Println(y)
	}
	// out scope
	//fmt.Println(y)
	fmt.Println(increment())
	fmt.Println(increment())

	x = 0
	increments := func() int {
		x++
		return x
	}
	fmt.Println(increments())
	fmt.Println(increments())

}

func order_matters() {
	// fmt.Println(x)
	fmt.Println(y)
	// x := 42
	// x declared but not used
}

var y = 42

func max(x int) int {
	return 42 + x
}

func variable_shadowing() {
	max := max(7)
	fmt.Println(max)
}

func main() {
	package_scope()
	block_scope()
	order_matters()
	variable_shadowing()
}
