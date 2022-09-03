package basic

import "testing"

var result string

func BenchmarkConcatenateBuffer(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = ConcatenateBuffer("test2", "test3")
	}
	result = s
}

func BenchmarkConcatenateJoin(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = ConcatenateJoin("test2", "test3")
	}
	result = s
}
