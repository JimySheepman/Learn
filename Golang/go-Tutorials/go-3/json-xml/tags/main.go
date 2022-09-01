package main

import (
	"fmt"
	"reflect"
)

// Here the field name is an empty string. We can
// add the omitempty directive into the struct tag
// to make it disappear from the outputted JSON
type Product struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"-"`
}

func main() {
	p := Product{ID: 32}
	t := reflect.TypeOf(p)
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("field Name : %s\n", t.Field(i).Name)
		fmt.Printf("field Tag : %s\n", t.Field(i).Tag)
	}
}
