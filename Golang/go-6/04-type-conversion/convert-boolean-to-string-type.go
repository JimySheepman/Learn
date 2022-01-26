package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var b bool = true
	fmt.Println(reflect.TypeOf(b))

	var s string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))

	bo := true
	st := fmt.Sprintf("%v", bo)
	fmt.Println(st)
	fmt.Println(reflect.TypeOf(s))
}
