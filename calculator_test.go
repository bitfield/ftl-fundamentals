package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a      float64
	b      float64
	want   float64
	tcName string
}

func TestAdd(t *testing.T) {
	t.Parallel()
	tCases := []testCase{
		{2, 2.030000, 4.030000, "some fractions1"},
		{6, 3.99, 9.99, "some fractions2"},
		{2, 2.03675, 4.03675, "some fractions3"},
	}
	for _, tc := range tCases {
		if got := calculator.Add(tc.a, tc.b); tc.want != got {
			t.Errorf("want %f, got %f :: Test: %s", tc.want, got, tc.tcName)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(2, 4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 8
	got := calculator.Multiply(2, 4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	tCases := []testCase{
		{2.030000, 20, 9.852216, "some fractions1"},
		{0, 3.99, 0, "some fractions2"},
		{2, 0, 0, "some fractions3"},
	}
	for _, tc := range tCases {
		if got := calculator.Add(tc.a, tc.b); tc.want != got {
			t.Errorf("want %f, got %f :: Test: %s", tc.want, got, tc.tcName)
		}
	}
}
