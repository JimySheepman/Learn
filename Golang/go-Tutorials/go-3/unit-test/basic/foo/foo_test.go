package foo

import "testing"

func TestFoo(t *testing.T) {
	expected := "Foo"
	actual := Foo()
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}
