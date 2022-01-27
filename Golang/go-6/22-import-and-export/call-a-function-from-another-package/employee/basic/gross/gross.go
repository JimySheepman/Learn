package gross

import (
	b "employee/basic"
)

func GrossSalary() int {
	a, t := b.Calculation()
	return ((b.Basic + a) - t)
}
