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
			t.Errorf("want %f, got %f", want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Multiply(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
