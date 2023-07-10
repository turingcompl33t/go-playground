package table

import (
	"errors"
	"fmt"
)

func DoMath(x int, y int, op string) (int, error) {
	switch op {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		if y == 0 {
			return 0, errors.New("division by zero")
		}
		return x / y, nil
	default:
		return 0, fmt.Errorf("unknown operator %s", op)
	}
}
