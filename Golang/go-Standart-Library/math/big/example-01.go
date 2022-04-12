package main

import (
	"fmt"
	"math/big"
)

func recur(n, lim int64) *big.Rat {
	term := new(big.Rat)
	if n%3 != 1 {
		term.SetInt64(1)
	} else {
		term.SetInt64((n - 1) / 3 * 2)
	}

	if n > lim {
		return term
	}

	frac := new(big.Rat).Inv(recur(n+1, lim))

	return term.Add(term, frac)
}
func main() {
	for i := 1; i <= 15; i++ {
		r := recur(0, int64(i))

		fmt.Printf("%-13s = %s\n", r, r.FloatString(8))
	}
}
