// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplication of both
// the numbers
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing the first number
// by second number. It returns error if the division operation is not valid
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divide by Zero")
	}
	return a / b, nil
}
