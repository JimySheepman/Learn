package main

import (
	"fmt"
	"testing"
)

func Sum(n int64) int64 {
	return n * (n + 1) / 2
}

func BenchmarkSum(b *testing.B) {
	fmt.Println("b.N:", b.N)
	for i := 0; i < b.N; i++ {
		Sum(1000000)
	}
}

func BenchmarkSumParallel(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Sum(1000000)
		}
	})
}
