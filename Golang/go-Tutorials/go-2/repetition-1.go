package main

import "fmt"

func main() {
	var i, j, k int
	var c, ch byte
	var f, salary float32
	d := 42
	var x float64
	x = 20.0
	var str string = "Merlins"
	fmt.Println(i, j, k, c, ch, f, salary, d, x, str)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("d is of type %T\n", d)
	fmt.Printf("str is of type %T\n", str)

	fmt.Printf("Hello\tWorld!\n")
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("value of area : %d\n", area)
}
