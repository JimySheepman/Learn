package main

import "fmt"

func main() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	fmt.Println(employee) // Initial Map

	employee["Rocky"] = 30 // Add element
	employee["Josef"] = 40

	fmt.Println(employee)
}
