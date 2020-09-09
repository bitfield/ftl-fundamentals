// Package calculator provides a library for simple calculations in Go.
package calculator

import "fmt"

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

//Multiply takes two numbers and returns the result of multiplying them by each other.
func Multiply(a, b float64) float64 {
	return a * b
}

//Divide takes two numbers and returns the result of dividing the first by the second.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Dividing by zero is not valid: %f / %f", a, b)
	}

	return a / b, nil
}

//Sqrt takes a number and returns the square root.
func Sqrt(a float64) (float64, error) {
	return 0, nil
}
