package main

import "fmt"

func main() {
	var ex01 bool = true

	var ex02 string = "Gopher"

	var ex03 rune = 'G'

	var ex04 uintptr = 1
	var ex05 uint8 = 1
	var ex06 uint16 = 1
	var ex07 uint = 1
	var ex08 uint32 = 1
	var ex09 uint64 = 1

	var ex10 int8 = 1
	var ex11 int16 = 1
	var ex12 int = 1
	var ex13 int32 = 1
	var ex14 int64 = 1

	var ex15 float32 = 1.0
	var ex16 float64 = 1.0

	var ex17 complex64 = 1
	var ex18 complex128 = 1

	var ex19 byte = 1

	var ex20 [2]int

	var ex21 *int

	var ex22 func(name, firstname string) string
	_ = ex22

	var ex23 []int

	var ex24 map[string]int

	var ex25 chan<- bool

	type ex26 struct {
		Name     string
		Capacity uint8
		Rooms    uint8
		Smoking  bool
	}

	type Firstname string
	type Currency string
	type VATRate float64

	fmt.Println(
		ex01, ex02, ex03, ex04, ex05,
		ex06, ex07, ex08, ex09, ex10,
		ex11, ex12, ex13, ex14, ex15,
		ex16, ex17, ex18, ex19, ex20,
		ex21, ex23, ex24, ex25,
	)
}
