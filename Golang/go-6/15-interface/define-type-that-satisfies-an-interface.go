package main

import "fmt"

type Employee interface {
	PrintName(name string)
	PrintSalary(basic int, tax int) int
}

type Emp int

func (e Emp) PrintName(name string) {
	fmt.Println("Emloyee Id:\t", e)
	fmt.Println("Employee Name:\t", name)
}

func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

func main() {
	var e1 Employee
	e1 = Emp(1)
	e1.PrintName("John Doe")
	fmt.Println("Employee Salary:", e1.PrintSalary(25000, 5))
}
