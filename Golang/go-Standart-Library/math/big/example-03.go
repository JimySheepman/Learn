package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	const prec = 200

	steps := int(math.Log2(prec))

	two := new(big.Float).SetPrec(prec).SetInt64(2)
	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)

	x := new(big.Float).SetPrec(prec).SetInt64(1)

	t := new(big.Float)

	for i := 0; i <= steps; i++ {
		t.Quo(two, x)
		t.Add(x, t)
		x.Mul(half, t)
	}

	fmt.Printf("sqrt(2) = %.50f\n", x)
	t.Mul(x, x)
	fmt.Printf("error = %e\n", t.Sub(two, t))
}
