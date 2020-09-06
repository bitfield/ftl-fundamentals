// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers or more and returns the result of adding them together.
func Add(a, b float64, numbers ...float64) float64 {
	sum := a + b
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// Subtract takes two numbers or more and returns the result of subtracting them
// from the first.
func Subtract(a, b float64, numbers ...float64) float64 {
	result := a - b
	for _, n := range numbers {
		result -= n
	}
	return result
}

// Multiply takes two or more numbers and returns the result of multiplying them together.
func Multiply(a, b float64, numbers ...float64) float64 {
	result := a*b
	for _, n := range numbers {
		result *= n
	}
	return result
}

// Divide takes two number and returns the result of dividing the first
// with the second. If the second number is zero it returns an error.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by 0")
	}
	return a / b, nil
}

// Sqrt takes one number and returns its square root.
// If the number is negative it returns an error.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("negative number")
	}
	return math.Sqrt(a), nil
}
