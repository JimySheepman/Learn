package main

import (
	"fmt"
	"reflect"
)

func main() {

	myFunc := func() int {
		fmt.Println("I am a func literal")
		return 42
	}

	fmt.Println(reflect.TypeOf(myFunc))

	funcValue := func() int {
		fmt.Println("I am a func literal invoked")
		return 42

	}()
	fmt.Println(reflect.TypeOf(funcValue))
}
