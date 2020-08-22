// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
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

// Multiply takes two numbers and returns the result of multiplying them
// together.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing the first by the
// second, or an error if the second value is zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("bad inputs %f, %f: division by zero is not allowed", a, b)
	}
	return a / b, nil
}

// Sqrt takes a number and returns its square root, or an error if the value is
// negative.
func Sqrt(input float64) (float64, error) {
	if input < 0 {
		return 0, fmt.Errorf("bad input %f: square root of a negative number is not allowed", input)
	}
	return math.Sqrt(input), nil
}
