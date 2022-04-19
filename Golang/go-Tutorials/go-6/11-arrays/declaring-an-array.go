package main

import (
	"fmt"
	"reflect"
)

func main() {
	var intArray [5]int
	var strArray [5]string

	fmt.Println(reflect.ValueOf(intArray).Kind())
	fmt.Println(reflect.ValueOf(strArray).Kind())
}
