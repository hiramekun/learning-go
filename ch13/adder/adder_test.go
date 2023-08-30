package adder

import "testing"

func Test_addNumbers(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Error("addNumbers(2, 3) should be 5, but doesn't match", result)
	}
}
