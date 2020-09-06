package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
)

type TestCase struct {
	name        string
	a, b        float64
	want        float64
	errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{name: "Two positive numbers", a: 2, b: 2, want: 4},
		{name: "One negative number", a: -1, b: 0, want: -1},
		{name: "One zero number", a: 0, b: 0, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s - Add(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < 100; i++ {
		a, b := r.Float64(), r.Float64()
		want := a + b
		got := calculator.Add(a, b)
		if want != got {
			t.Fatalf("Random Add(%f, %f): want %f, got %f", a, b, want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{name: "Two positive numbers", a: 2, b: 1, want: 1},
		{name: "Two positive numbres which substract to a negative", a: 2, b: 3, want: -1},
		{name: "Two negative numbers which substract to a positive", a: -1, b: -3, want: 2},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s - Substract(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{name: "Two positive numbers", a: 2, b: 3, want: 6},
		{name: "One zero number", a: 0, b: 4, want: 0},
		{name: "One negative number", a: 3, b: -1, want: -3},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s - Multiply(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{name: "Two positive numbers", a: 4, b: 2, want: 2, errExpected: false},
		{name: "Division by 0", a: 8, b: 0, want: 0, errExpected: true},
		{name: "0 divide by a number", a: 0, b: 1, want: 0, errExpected: false},
		{name: "Negative number", a: -3, b: 1, want: -3, errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("%s - Divide(%f, %f): errExpected %t, error %v", tc.name, tc.a, tc.b, tc.errExpected, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("%s - Divide(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{name: "Positive Number", a: 144, want: 12, errExpected: false},
		{name: "Negative number", a: -2, want: 0, errExpected: true},
		{name: "Zero", a: 0, want: 0, errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("%s - Sqrt(%f): errExpected %t, error %v", tc.name, tc.a, tc.errExpected, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Fatalf("%s - Sqrt(%f): want %f, got %f", tc.name, tc.a, tc.want, got)
		}
	}
}
