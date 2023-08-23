package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {
	testAdd := 4
	add, _ := Calc("+", 2, 2)

	testSubtract := 0
	subtract, _ := Calc("-", 2, 2)

	testMultiple := 4
	multiply, _ := Calc("*", 2, 2)

	testDivide := 1
	divide, _ := Calc("/", 2, 2)

	_, err := Calc("/", 2, 0)
	testError := err.Error()

	assert.Equal(t, testAdd, add)
	assert.Equal(t, testSubtract, subtract)
	assert.Equal(t, testMultiple, multiply)
	assert.Equal(t, testDivide, divide)
	assert.EqualError(t, err, testError)
}
