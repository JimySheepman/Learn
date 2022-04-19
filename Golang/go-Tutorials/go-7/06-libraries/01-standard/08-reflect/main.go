package main

import (
	"fmt"
	"reflect"
)

func main() {
	Sample("Type", Type)
	Sample("TypeOf", TypeOf)
}

func Type() {

}

func TypeOf() {
	fn := func(value interface{}) {
		fmt.Printf("%#v -> %s\n", value, reflect.TypeOf(value))
	}

	type Post struct {
		Title string
	}

	fn("Hello")
	fn(100)
	fn(7.2)
	fn(true)
	fn(struct{}{})
	fn(Post{"Post 1"})
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
