package main

import (
	"fmt"
	"math/big"
)

func main() {
	i := new(big.Int)
	i.SetString("644", 8)
	fmt.Println(i)
}
