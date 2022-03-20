package week_01

import "testing"

func TestIsEmpty(t *testing.T) {
	table := []interface{}{false, "", 0, nil}
	for _, v := range table {
		expexted := IsEmpty(v)
		if expexted != true {
			t.Error("IsEmpty of \"", v, "\" is failed.")
		}
	}
}
