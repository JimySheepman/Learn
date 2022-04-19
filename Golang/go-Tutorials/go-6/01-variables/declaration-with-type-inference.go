package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i = 10
	var s = "Canada"

	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(s))
}
