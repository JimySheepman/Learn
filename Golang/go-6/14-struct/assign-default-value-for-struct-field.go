package main

import "fmt"

type Employee struct {
	Name string
	Age  int
}

func (obj *Employee) Info() {
	if obj.Name == "" {
		obj.Name = "John Doe"
	}
	if obj.Age == 0 {
		obj.Age = 25
	}
}

func main() {
	emp1 := Employee{Name: "Mr. Fred"}
	emp1.Info()
	fmt.Println(emp1)

	emp2 := Employee{Age: 26}
	emp2.Info()
	fmt.Println(emp2)
}
