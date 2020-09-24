// Package calculator provides a library for simple calculations in Go.
package calculator

import "errors"

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return b - a
}

// Multiply takes two numbers and returns the result of multplying the second
// by the first.
func Multiply(a, b float64) float64 {
	return b * a
}

// Divide takes two numbers and returns the result of dividing the second
// by the first.
func Divide(a, b float64) (float64, error) {
	if a != 0 {
		return b / a, nil
	}
	return 0, errors.New("math: Divide by Zero")
}
