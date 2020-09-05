package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}


func almostEqual(a, b float64) bool {
	const float64EqualityThreshold = 1e-9
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var successTestCases = []struct {
		name string
		want float64
		a    float64
		b    float64
	}{
		{
			name: "Two positive numbers",
			want: 3,
			a:    5,
			b:    2,
		},
		{
			name: "Zero subtracted with any number",
			want: -2,
			a:    0,
			b:    2,
		},
		{
			name: "Positive and negative number",
			want: -9,
			a:    -5,
			b:    4,
		},
		{
			name: "Two negative numbers",
			want: -5,
			a:    -8,
			b:    -3,
		},
		{
			name: "Two fractional numbers",
			want: 0.1,
			a:    3.3,
			b:    3.2,
		},
	}

	for _, test := range successTestCases {
		got := calculator.Subtract(test.a, test.b)
		if !almostEqual(test.want, got) {
			t.Errorf("Test Subtract(%f, %f) %s failed: want %f, got %f",
				test.a, test.b, test.name, test.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var successTestCases = []struct {
		name string
		want float64
		a    float64
		b    float64
	}{
		{
			name: "Two positive numbers",
			want: 10,
			a:    5,
			b:    2,
		},
		{
			name: "Zero multiplied with any number",
			want: 0,
			a:    0,
			b:    2,
		},
		{
			name: "Positive and negative number",
			want: -20,
			a:    -5,
			b:    4,
		},
		{
			name: "Two negative numbers",
			want: 24,
			a:    -8,
			b:    -3,
		},
		{
			name: "Two fractional numbers",
			want: 4.03,
			a:    1.3,
			b:    3.1,
		},
	}

	for _, test := range successTestCases {
		got := calculator.Multiply(test.a, test.b)
		if !almostEqual(test.want, got) {
			t.Errorf("Test Multiple(%f, %f) %s failed: want %f, got %f",
				test.a, test.b, test.name, test.want, got)
		}
	}

}

func TestDivide(t *testing.T) {
	t.Parallel()
	var successTestCases = []struct {
		name        string
		want        float64
		a           float64
		b           float64
		errExpected bool
	}{
		{
			name:        "Two positive numbers",
			want:        5,
			a:           10,
			b:           2,
			errExpected: false,
		},
		{
			name:        "Zero divided with any number",
			want:        0,
			a:           0,
			b:           2,
			errExpected: false,
		},
		{
			name:        "Positive and negative number",
			want:        -1.25,
			a:           -5,
			b:           4,
			errExpected: false,
		},
		{
			name:        "Two negative numbers",
			want:        4,
			a:           -8,
			b:           -2,
			errExpected: false,
		},
		{
			name:        "Two fractional numbers",
			want:        2.75,
			a:           3.3,
			b:           1.2,
			errExpected: false,
		},
		{
			name:        "Divide by zero",
			want:        0,
			a:           2,
			b:           0,
			errExpected: true,
		},
	}

	for _, test := range successTestCases {
		got, err := calculator.Divide(test.a, test.b)
		if test.errExpected {
			if err == nil {
				t.Errorf("Test Divide (%f, %f) %s failed: Expected error got nil",
					test.a, test.b, test.name)
			}
		} else {
			if err != nil {
				t.Errorf("Test Divide (%f, %f) %s failed: Got err expected nil",
					test.a, test.b, test.name)
			}
			if !almostEqual(test.want, got) {
				t.Errorf("Test Divide (%f, %f) %s failed: want %f, got %f",
					test.a, test.b, test.name, test.want, got)
			}
		}

	}

}
