package main

import "fmt"

func main() {
	var employee = map[string]int{"mark": 10, "sandy": 20}
	fmt.Println(employee)

	employee["mark"] = 50
	fmt.Println(employee)
}
