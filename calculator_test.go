package calculator_test

import (
	"calculator"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

type testCase struct {
	name        string
	firstInput  float64
	secondInput float64
	expected    float64
}

type testCaseWithErr struct {
	name        string
	firstInput  float64
	secondInput float64
	errExpected bool
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
		want := c.expected
		got := f(c.firstInput, c.secondInput)
		if want != got {
			t.Errorf(c.name+": want %f, got %f", want, got)
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
		{"Subtract a negative number from a positive number", 5, -3, 8},
		{"Subtract a postitive number from a negative number", -4, 2, -6},
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
		{"Multiply one postive number by one negative number", -5, 3, -15},
	}

	t.Parallel()
	for _, c := range cases {
		t.Run(c.name, testFunc(c, calculator.Multiply))
	}
}

func TestDivide(t *testing.T) {
	cases := []testCaseWithErr{
		{"Divide two positive numbers", 4, 2, false, 2},
		{"Divide two negative numbers", -4, -2, false, 2},
		{"Divide one positive number by one negative number", 4, -2, false, -2},
		{"Divide one negative number by one postive number", -4, 2, false, -2},
		{"Divide two fractional positive numbers", 4.2, 2.1, false, 2},
		{"Divide two fractional negative numbers equaling a whole number", -4.2, -2.1, false, 2},
		{"Divide one positive number by zero", 4, 0, true, 0},
	}

	t.Parallel()
	for _, c := range cases {
		t.Run(c.name, testDivideFunc(c, calculator.Divide))
	}
}

func testDivideFunc(c testCaseWithErr, f func(float64, float64) (float64, error)) func(*testing.T) {
	return func(t *testing.T) {
		want := c.expected
		got, err := f(c.firstInput, c.secondInput)
		if err != nil && !c.errExpected {
			t.Errorf(c.name+": wanted %f got an error: %v ", want, err)
		}

		if want != got {
			t.Errorf(c.name+": want %f, got %f", want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	t.Parallel()
	for i := 0; i < 100; i++ {
		first := r.Float64() * 100
		second := r.Float64() * 100
		name := fmt.Sprintf("Adding %f and %f", first, second)
		result := first + second
		c := testCase{name, first, second, result}

		t.Run(c.name, testFunc(c, calculator.Add))
	}

}

func TestSqrt(t *testing.T) {
	cases := []testCaseWithErr{
		{"Square root of zero should be zero", 0, 0, false, 0},
		{"Square root of a negative number should return an error", -4, 0, true, 0},
		{"Square root of a positive whole number", 4, 0, false, 2},
	}

	t.Parallel()
	for _, c := range cases {
		t.Run(c.name, testSqrtFunc(c, calculator.Sqrt))
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < 100; i++ {
		value := r.Float64() * 100
		name := fmt.Sprintf("Square root of %f", value)
		result := math.Sqrt(value)
		c := testCaseWithErr{name, value, 0, false, result}

		t.Run(c.name, testSqrtFunc(c, calculator.Sqrt))
	}
}

func testSqrtFunc(c testCaseWithErr, f func(float64) (float64, error)) func(*testing.T) {
	return func(t *testing.T) {
		want := c.expected
		got, err := f(c.firstInput)
		if err != nil && !c.errExpected {
			t.Errorf(c.name+": wanted %f got an error: %v ", want, err)
		}

		if want != got {
			t.Errorf(c.name+": want %f, got %f", want, got)
		}
	}
}
