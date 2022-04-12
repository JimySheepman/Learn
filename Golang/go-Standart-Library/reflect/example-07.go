package main

import (
	"fmt"
	"reflect"
)

func main() {
	type user struct {
		firstName string
		lastName  string
	}

	type data struct {
		user
		firstName string
		lastName  string
	}

	u := data{
		user:      user{"Embedded John", "Embedded Doe"},
		firstName: "John",
		lastName:  "Doe",
	}

	s := reflect.ValueOf(u).FieldByIndex([]int{0, 1})
	fmt.Println("embedded last name:", s)
}
