package main

import (
	"fmt"
	"math/bits"
)

func main() {
	n1 := []uint64{0, 6}
	n2 := []uint64{0, 3}
	quo, rem := bits.Div64(n1[0], n1[1], n2[1])
	nsum := []uint64{quo, rem}
	fmt.Printf("[%v %v] / %v = %v\n", n1[0], n1[1], n2[1], nsum)

	n1 = []uint64{2, 0x8000000000000000}
	n2 = []uint64{0, 0x8000000000000000}
	quo, rem = bits.Div64(n1[0], n1[1], n2[1])
	nsum = []uint64{quo, rem}
	fmt.Printf("[%v %v] / %v = %v\n", n1[0], n1[1], n2[1], nsum)
}
