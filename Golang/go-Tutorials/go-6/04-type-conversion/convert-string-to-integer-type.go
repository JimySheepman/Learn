package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	strVar := "100"
	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar2, err := strconv.ParseInt(strVar, 0, 8)
	fmt.Println(intVar2, err, reflect.TypeOf(intVar2))

	intVar2, err = strconv.ParseInt(strVar, 0, 16)
	fmt.Println(intVar2, err, reflect.TypeOf(intVar2))

	intVar2, err = strconv.ParseInt(strVar, 0, 32)
	fmt.Println(intVar2, err, reflect.TypeOf(intVar2))

	intVar2, err = strconv.ParseInt(strVar, 0, 64)
	fmt.Println(intVar2, err, reflect.TypeOf(intVar2))

	intValue := 0
	t, err := fmt.Sscan(strVar, &intValue)
	fmt.Println(intValue, err, reflect.TypeOf(intValue), t)
}
