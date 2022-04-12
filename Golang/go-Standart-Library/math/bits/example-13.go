package main

import (
	"fmt"
	"math/bits"
)

func main() {
	n1 := []uint32{0, 12}
	n2 := []uint32{0, 12}
	hi, lo := bits.Mul32(n1[1], n2[1])
	nsum := []uint32{hi, lo}
	fmt.Printf("%v * %v = %v\n", n1[1], n2[1], nsum)

	n1 = []uint32{0, 0x80000000}
	n2 = []uint32{0, 2}
	hi, lo = bits.Mul32(n1[1], n2[1])
	nsum = []uint32{hi, lo}
	fmt.Printf("%v * %v = %v\n", n1[1], n2[1], nsum)
}
