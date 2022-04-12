package main

import (
	"fmt"
	"math/bits"
)

func main() {
	n1 := []uint64{33, 23}
	n2 := []uint64{21, 12}
	d1, carry := bits.Sub64(n1[1], n2[1], 0)
	d0, _ := bits.Sub64(n1[0], n2[0], carry)
	nsum := []uint64{d0, d1}
	fmt.Printf("%v - %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)

	n1 = []uint64{3, 0x7fffffffffffffff}
	n2 = []uint64{1, 0x8000000000000000}
	d1, carry = bits.Sub64(n1[1], n2[1], 0)
	d0, _ = bits.Sub64(n1[0], n2[0], carry)
	nsum = []uint64{d0, d1}
	fmt.Printf("%v - %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)
}
