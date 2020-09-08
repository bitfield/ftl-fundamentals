package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	name        string
	firstInput  float64
	secondInput float64
	expected    float64
}

func TestAdd(t *testing.T) {
	cases := []testCase{
		{"Add two positive numbers", 2, 2, 4},
		{"Add two negative numbers", -2, -2, -4},
		{"Add one negative and one positive number equaling zero", -1, 1, 0},
		{"Add one fractional and one whole number", 5.4, 2, 7.4},
		{"Add two fractional numbers", 2.3, 4.3, 6.6},
		{"Add two fractional numbers equaling a whole number", 2.3, 3.7, 6},
		{"Add two negative fractional numbers", -1.5, -2.5, -4},
	}

	t.Parallel()
	for _, c := range cases {
		t.Run(c.name, testFunc(c, calculator.Add))
	}
}

func testFunc(c testCase, f func(float64, float64) float64) func(*testing.T) {
	return func(t *testing.T) {
		var want float64 = c.expected
		got := f(c.firstInput, c.secondInput)
		if want != got {
			t.Errorf(c.name+": want %f of type %T, got %f of type %T", want, want, got, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	cases := []testCase{
		{"Subtract two positive numbers", 5, 1, 4},
		{"Subtract two negative numbers", -5, -1, -4},
		{"Subtract two negative numbers equaling zero", -1, -1, 0},
		{"Subtract one fractional number from one whole number", 5.0, 1.5, 3.5},
		{"Subtract two fractional numbers", 2.3, 4.3, -2},
		{"Subtract two fractional numbers equaling a whole number", 2.5, 1.5, 1},
		{"Subtract two negative fractional numbers", -2.5, -1.5, -1},
	}

	t.Parallel()
	for _, c := range cases {
		t.Run(c.name, testFunc(c, calculator.Subtract))
	}
}

func TestMultiply(t *testing.T) {
	cases := []testCase{
		{"Multiply two positive numbers", 2, 2, 4},
		{"Multiply two negative numbers", -2, -2, 4},
		{"Multiply one postive number by zero", 1, 0, 0},
		{"Multiply one fractional by one whole number", 5.4, 2, 10.8},
		{"Multiply two fractional numbers", 2.6, 5.3, 13.78},
		{"Multiply two fractional numbers equaling a whole number", 1.5, 4, 6},
		{"Multiply two negative fractional numbers", -1.5, -2.5, 3.75},
	}

	t.Parallel()
	for _, c := range cases {
		t.Run(c.name, testFunc(c, calculator.Multiply))
	}
}
