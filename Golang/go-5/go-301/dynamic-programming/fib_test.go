package main

import "testing"

func BenchmarkFib1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Fib1(15)
	}
}

func BenchmarkFib2(b *testing.B) {
	n := 15
	memo := make([]int, n)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Fib2(n, memo)
	}
}

func BenchmarkFib3(b *testing.B) {
	n := 15
	memo := make([]int, n+1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Fib3(n, memo)
	}
}
