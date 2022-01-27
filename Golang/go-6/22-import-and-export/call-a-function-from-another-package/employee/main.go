package main

import (
	b "employee/basic"
	"employee/basic/gross"

	"fmt"
)

func main() {
	b.Basic = 10000
	fmt.Println(gross.GrossSalary())
}
