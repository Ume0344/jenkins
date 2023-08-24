package calculator

import (
	"errors"
	"fmt"
)

// Calc is a basic Calculator
// Param x: a number of int type
// Param y: a number of int type
// returns z, err: result of requested operation and error if it occurs.
func Calc(operator string, x int, y int) (int, error) {
	var z int

	switch operator {
	case "+":
		z = x + y
	case "-":
		z = x - y
	case "*":
		z = x * y
	case "/":
		if y != 0 {
			z = x / y
		} else {
			return 0, errors.New("Second number should not be 0 for division\n")
		}

	default:
		fmt.Print("Operator not from {+, -, *, /}\n")
	}

	return z, nil
}
