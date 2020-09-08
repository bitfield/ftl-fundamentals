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
		{"Two positive numbers", 2, 2, 4},
		{"Two negative numbers", -2, -2, -4},
		{"One negative and one positive number equaling zero", -1, 1, 0},
		{"One fractional and one whole number", 5.4, 2, 7.4},
		{"Two fractional numbers", 2.3, 4.3, 6.6},
		{"Two fractional numbers equaling a whole number", 2.3, 3.7, 6},
		{"Two negative fractional numbers", -1.5, -2.5, -4},
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
			t.Errorf("want %f of type %T, got %f of type %T", want, want, got, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	cases := []testCase{
		{"Two positive numbers", 5, 1, 4},
		{"Two negative numbers", -5, -1, -4},
		{"Two negative numbers equaling zero", -1, -1, 0},
		{"One fractional and one whole number", 5.5, 3.0, 2.5},
		{"Two fractional numbers", 2.3, 4.3, -2},
		{"Two fractional numbers equaling a whole number", 2.5, 1.5, 1},
		{"Two negative fractional numbers", -2.5, -1.5, -1},
	}

	t.Parallel()
	for _, c := range cases {
		t.Run(c.name, testFunc(c, calculator.Subtract))
	}
}

func TestMultiply(t *testing.T) {
	cases := []testCase{
		{"Two positive numbers", 2, 2, 4},
		{"Two negative numbers", -2, -2, 4},
		{"One postive number and zero", 1, 0, 0},
		{"One fractional and one whole number", 5.4, 2, 10.8},
		{"Two fractional numbers", 2.6, 5.3, 13.78},
		{"Two fractional numbers equaling a whole number", 1.5, 4, 6},
		{"Two negative fractional numbers", -1.5, -2.5, 3.75},
	}

	t.Parallel()
	for _, c := range cases {
		t.Run(c.name, testFunc(c, calculator.Multiply))
	}
}
