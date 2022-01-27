package main

import (
	"fmt"
	"reflect"
)

type rectangle struct {
	length  float64
	breadth float64
	color   string
}

func main() {
	var rect1 = rectangle{10, 20, "Green"}
	fmt.Println(reflect.TypeOf(rect1))         // main.rectangle
	fmt.Println(reflect.ValueOf(rect1).Kind()) // struct

	rect2 := rectangle{length: 10, breadth: 20, color: "Green"}
	fmt.Println(reflect.TypeOf(rect2))         // main.rectangle
	fmt.Println(reflect.ValueOf(rect2).Kind()) // struct

	rect3 := new(rectangle)
	fmt.Println(reflect.TypeOf(rect3))         // *main.rectangle
	fmt.Println(reflect.ValueOf(rect3).Kind()) // ptr

	var rect4 = &rectangle{}
	fmt.Println(reflect.TypeOf(rect4))         // *main.rectangle
	fmt.Println(reflect.ValueOf(rect4).Kind()) // ptr
}
