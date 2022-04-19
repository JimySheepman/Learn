package main

import (
	"fmt"
	"internal-sample/utils"
	"internal-sample/utils/calc"
)

func main() {
	a := utils.Rand(0, 10)
	b := utils.Rand(0, 10)
	sum := calc.Sum(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, sum)
}
