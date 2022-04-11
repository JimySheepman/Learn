package main

import (
	"sort"
	"testing"
)

func main() {
	testing.Benchmark(func(b *testing.B) {
		var compares int64
		for i := 0; i < b.N; i++ {
			s := []int{5, 4, 3, 2, 1}
			sort.Slice(s, func(i, j int) bool {
				compares++
				return s[i] < s[j]
			})
		}
		b.ReportMetric(float64(compares)/float64(b.N), "compares/op")
	})
}
