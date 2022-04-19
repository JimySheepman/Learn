package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var f float64 = 3.1415926535
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(f)

	var s string = strconv.FormatFloat(f, 'E', -1, 32)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s)

	b := 13.454
	fmt.Println(reflect.TypeOf(b))

	s = fmt.Sprintf("%v", b)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}
