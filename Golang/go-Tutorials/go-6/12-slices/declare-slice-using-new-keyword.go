package main

import (
	"fmt"
	"reflect"
)

func main() {
	var intSlice = new([50]int)[0:10]

	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
}
