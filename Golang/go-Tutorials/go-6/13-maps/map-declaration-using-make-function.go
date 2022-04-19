package main

import "fmt"

func main() {
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20
	fmt.Println(employee)

	employeeList := make(map[string]int)
	employeeList["Mark"] = 10
	employeeList["Sandy"] = 20
	fmt.Println(employeeList)

	employeeList = make(map[string]int)

	fmt.Println(len(employee))
	fmt.Println(len(employeeList))
}
