package main

import (
	"fmt"
	"math/bits"
)

func main() {
	n1 := []uint64{0, 12}
	n2 := []uint64{0, 12}
	hi, lo := bits.Mul64(n1[1], n2[1])
	nsum := []uint64{hi, lo}
	fmt.Printf("%v * %v = %v\n", n1[1], n2[1], nsum)

	n1 = []uint64{0, 0x8000000000000000}
	n2 = []uint64{0, 2}
	hi, lo = bits.Mul64(n1[1], n2[1])
	nsum = []uint64{hi, lo}
	fmt.Printf("%v * %v = %v\n", n1[1], n2[1], nsum)
}
