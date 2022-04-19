package main

import "fmt"

func main() {
	var employee = make(map[string]int)
	employee["mark"] = 10
	employee["sandy"] = 20
	employee["rocky"] = 30
	employee["josef"] = 40

	fmt.Println(employee)

	delete(employee, "mark")
	fmt.Println(employee)

}
