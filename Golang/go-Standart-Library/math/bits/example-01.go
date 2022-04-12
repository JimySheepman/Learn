package main

import (
	"fmt"
	"math/bits"
)

func main() {
	n1 := []uint32{33, 12}

	n2 := []uint32{21, 23}

	d1, carry := bits.Add32(n1[1], n2[1], 0)
	d0, _ := bits.Add32(n1[0], n2[0], carry)
	nsum := []uint32{d0, d1}
	fmt.Printf("%v + %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)

	n1 = []uint32{1, 0x80000000}
	n2 = []uint32{1, 0x80000000}
	d1, carry = bits.Add32(n1[1], n2[1], 0)
	d0, _ = bits.Add32(n1[0], n2[0], carry)
	nsum = []uint32{d0, d1}
	fmt.Printf("%v + %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)
}
