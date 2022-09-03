package main

import (
	"fmt"
	"testing"
)

func main() {
	res := testing.Benchmark(BenchmarkConcatenateBuffer)
	fmt.Printf("Memory allocations : %d \n", res.MemAllocs)
	fmt.Printf("Number of bytes allocated: %d \n", res.Bytes)
	fmt.Printf("Number of run: %d \n", res.N)
	fmt.Printf("Time taken: %s \n", res.T)
}

// ..
func BenchmarkConcatenateBuffer(b *testing.B) {
	//..
}
