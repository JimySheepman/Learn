package main

import (
	"fmt"
	"math/bits"
)

func main() {
	n1 := []uint32{0, 6}
	n2 := []uint32{0, 3}
	quo, rem := bits.Div32(n1[0], n1[1], n2[1])
	nsum := []uint32{quo, rem}
	fmt.Printf("[%v %v] / %v = %v\n", n1[0], n1[1], n2[1], nsum)

	n1 = []uint32{2, 0x80000000}
	n2 = []uint32{0, 0x80000000}
	quo, rem = bits.Div32(n1[0], n1[1], n2[1])
	nsum = []uint32{quo, rem}
	fmt.Printf("[%v %v] / %v = %v\n", n1[0], n1[1], n2[1], nsum)
}
