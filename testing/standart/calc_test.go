package main

import (
	"fmt"
	"testing"
)

func TestSum_Simple(t *testing.T) {
	result := Sum(5, 7)
	expected := 12

	if result != expected {
		t.Error("Expected:", expected, "Got:", result)
	}
}

func TestSum_Cases(t *testing.T) {
	type testCase struct{ a, b, sum int }

	cases := []testCase{
		{5, 7, 12},
		{7, 12, 19},
		{3, 5, 8},
	}

	for _, c := range cases {
		result := Sum(c.a, c.b)
		expected := c.sum

		if result != expected {
			t.Error("Expected:", expected, "Got:", result)
		}
	}
}

func TestSum_SubTest(t *testing.T) {
	type testCase struct{ a, b, sum int }

	cases := []testCase{
		{5, 7, 12},
		{7, 12, 19},
		{3, 5, 8},
	}

	for _, c := range cases {
		name := fmt.Sprintf("%d + %d", c.a, c.b)

		t.Run(name, func(t *testing.T) {
			result := Sum(c.a, c.b)
			expected := c.sum

			if result != expected {
				t.Error("Expected:", expected, "Got:", result)
			}
		})
	}
}

// sudo go test -bench=Sum
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(4, 6)
	}
}

// this is for docs
func ExampleSum() {
	fmt.Println(Sum(4, 6))
	// output: 10
}
