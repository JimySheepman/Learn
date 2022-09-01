package main

import "fmt"

func main() {
	innerSlice1 := []float64{2.30, 4.01, 6.99, 8}
	innerSlice2 := []float64{1, 10.20, 30, 1.34, 23}
	my2DSlice := [][]float64{innerSlice1, innerSlice2}
	fmt.Println(my2DSlice)
}
