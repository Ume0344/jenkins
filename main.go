package main

import (
	"fmt"
	"jenkins/calculator"
)

func main() {
	x := 2
	y := 0
	z, err := calculator.Calc("/", x, y)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

	fmt.Printf("Result of requested operation is %d \n", z)
}
