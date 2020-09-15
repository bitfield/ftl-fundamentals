package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a, b float64
		want float64
	}
	testCases := []testCase{
		{name: "Two and two make four", a: 2, b: 2, want: 4},
		{name: "One and one make two", a: 1, b: 1, want: 2},
		{name: "N plus zero is N", a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Add(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a + b
		got := calculator.Add(a, b)
		if want != got {
			t.Fatalf("Add(%f, %f): want %f, got %f", a, b, want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a, b float64
		want float64
	}
	testCases := []testCase{
		{name: "N minus N is zero", a: 2, b: 2, want: 0},
		{name: "A positive minus a larger positive is a negative", a: 1, b: 2, want: -1},
		{name: "N minus zero is N", a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Subtract(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a, b float64
		want float64
	}
	testCases := []testCase{
		{name: "Two twos are four", a: 2, b: 2, want: 4},
		{name: "One times N is N", a: 1, b: 2, want: 2},
		{name: "N times zero is zero", a: 5, b: 0, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Multiply(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name        string
		a, b        float64
		want        float64
		errExpected bool
	}
	testCases := []testCase{
		{name: "N divided by N is 1", a: 2, b: 2, want: 1, errExpected: false},
		{name: "One divided by two is one-half", a: 1, b: 2, want: 0.5, errExpected: false},
		{name: "Division by zero is not allowed", a: 5, b: 0, want: 999, errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("%s: Divide(%f, %f): unexpected error status: %v", tc.name, tc.a, tc.b, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("%s: Divide(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name        string
		input       float64
		want        float64
		errExpected bool
	}
	testCases := []testCase{
		{name: "Square root of 4 is 2", input: 4, want: 2, errExpected: false},
		{name: "Square root of one is one", input: 1, want: 1, errExpected: false},
		{name: "Square root of a negative number is not allowed (unless you're a physicist)", input: -1, want: 999, errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.input)
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("%s: Sqrt(%f): unexpected error status: %v", tc.name, tc.input, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("%s: Sqrt(%f): want %f, got %f", tc.name, tc.input, tc.want, got)
		}
	}
}
