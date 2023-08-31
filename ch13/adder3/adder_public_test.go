package adder_test

import (
	"test_examples/adder"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	result := adder.AddNumbers(2, 3)
	if result != 5 {
		t.Error("addNumbers(2, 3) should be 5, but doesn't match", result)
	}
}
