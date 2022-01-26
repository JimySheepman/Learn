package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var i int64 = 125
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(i)

	var s string = strconv.FormatInt(i, 10)
	fmt.Println(reflect.TypeOf(s))

	fmt.Println("Base 10 value of s:", s)

	s = strconv.FormatInt(i, 8)
	fmt.Println("Base 8 value of s:", s)

	s = strconv.FormatInt(i, 16)
	fmt.Println("Base 16 value of s:", s)

	s = strconv.FormatInt(i, 32)
	fmt.Println("Base 32 value of s:", s)

	b := 1225
	fmt.Println(reflect.TypeOf(b))

	s = fmt.Sprintf("%v", b)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}
