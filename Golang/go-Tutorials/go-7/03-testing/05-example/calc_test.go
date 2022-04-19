package calc

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(5, 7)
	expected := 12

	if result != expected {
		t.Error("Expected:", expected, "Got:", result)
	}
}

func ExampleSum() {
	fmt.Println(Sum(4, 6))
}
