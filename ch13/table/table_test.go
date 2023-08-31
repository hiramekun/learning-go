package table

import (
	"testing"
)

/*
➜ go test -v
=== RUN   TestDoMath
=== RUN   TestDoMath/addition
=== RUN   TestDoMath/subtraction
=== RUN   TestDoMath/multiplication
=== RUN   TestDoMath/division
=== RUN   TestDoMath/division_by_zero
=== RUN   TestDoMath/invalid_operator
--- PASS: TestDoMath (0.00s)

	--- PASS: TestDoMath/addition (0.00s)
	--- PASS: TestDoMath/subtraction (0.00s)
	--- PASS: TestDoMath/multiplication (0.00s)
	--- PASS: TestDoMath/division (0.00s)
	--- PASS: TestDoMath/division_by_zero (0.00s)
	--- PASS: TestDoMath/invalid_operator (0.00s)

PASS
ok      test_examples/table     0.305s
*/
func TestDoMath(t *testing.T) {
	data := []struct {
		name     string
		num1     int
		num2     int
		op       string
		expected int
		errMsg   string
	}{
		{"addition", 1, 2, "+", 3, ""},
		{"subtraction", 1, 2, "-", -1, ""},
		{"multiplication", 1, 2, "*", 2, ""},
		{"division", 2, 2, "/", 1, ""},
		{"division by zero", 2, 0, "/", 0, "cannot divide by zero"},
		{"invalid operator", 2, 0, "x", 0, "invalid operator x"},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("expected %d, got %d", d.expected, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("expected error message `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
