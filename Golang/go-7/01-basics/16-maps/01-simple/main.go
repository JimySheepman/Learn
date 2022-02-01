package main

import "fmt"

func main() {
	var dict map[string]string

	if dict == nil {
		fmt.Println("dict is nil")
	}

	fmt.Printf("dict['name'] = %#v\n", dict["name"])
	fmt.Printf("len(dict) = %d\n", len(dict))

	dict = map[string]string{
		"name": "Gopher",
	}

	dict["age"] = "1"

	fmt.Printf("dict['name'] = %#v\n", dict["name"])
	fmt.Printf("dict['age'] = %#v\n", dict["age"])
	fmt.Printf("len(dict) = %d\n", len(dict))
}
