package table

import "testing"

func TestDoMathTable(t *testing.T) {
	data := []struct {
		name     string
		x        int
		y        int
		op       string
		expected int
		errMsg   string
	}{
		{"addition", 2, 2, "+", 4, ""},
		{"subtraction", 2, 2, "-", 0, ""},
		{"multiplication", 2, 2, "*", 4, ""},
		{"division", 2, 2, "/", 1, ""},
		{"bad_division", 2, 0, "/", 0, "division by zero"},
		{"bad_op", 2, 2, "?", 0, "unknown operator ?"},
	}

	// Table tests are more powerful than builtin dynamic test support
	// in other languages in that the go unit test library can automatically
	// convert dynamic tests into "first-class" tests that populate in the
	// test report output
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := DoMath(d.x, d.y, d.op)
			if result != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error message `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
