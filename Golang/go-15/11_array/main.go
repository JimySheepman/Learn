package main

import "fmt"

func array_init() {
	var x [58]int
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	x[42] = 777
	fmt.Println(x[42])

	var y [58]string

	for i := 65; i <= 122; i++ {
		y[i-65] = string(i)
	}

	fmt.Println(y)
	fmt.Println(y[42])
}
func main() {
	array_init()
}
