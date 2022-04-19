package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	s := "3.1415826535"
	f, err := strconv.ParseFloat(s, 8)
	fmt.Println(f, err, reflect.TypeOf(f))

	s1 := "-3.141"
	f1, err := strconv.ParseFloat(s1, 8)
	fmt.Println(f1, err, reflect.TypeOf(f1))

	s2 := "A-3141X"
	f2, err := strconv.ParseFloat(s2, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f2, err, reflect.TypeOf(f2))
	}
}
