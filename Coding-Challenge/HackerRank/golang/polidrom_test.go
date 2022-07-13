package main

import (
	"reflect"
	"testing"
)

func TestCheckPolidrom(t *testing.T) {
	type Test struct {
		input string
		want  bool
	}

	tests := []Test{
		{input: "Anna", want: false},
		{input: "civic", want: true},
		{input: "kayak", want: true},
		{input: "madam", want: true},
		{input: "mom", want: true},
		{input: "noon", want: true},
		{input: "racecar", want: true},
		{input: "radar", want: true},
		{input: "redder", want: true},
		{input: "refer", want: true},
		{input: "repaper", want: true},
		{input: "taco cat", want: false},
		{input: "my gym", want: false},
		{input: "tops spot", want: true},
		{input: "step on no pets.", want: false},
		{input: "Eva, can I see bees in a cave?", want: false},
	}

	for _, test := range tests {
		got := CheckPolidrom(test.input)
		if !reflect.DeepEqual(test.want, got) {
			t.Fatalf("Test case: %v, expected: %v, got: %v", test.input, test.want, got)
		}
	}
}
