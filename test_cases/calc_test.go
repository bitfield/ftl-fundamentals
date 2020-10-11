package calc_test

import (
	"calculator"
	"testing"
)



type testCase struct {
	a, b float64
	want float64
	errExpected bool
}


// A slice of test cases

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f,%f): want %f, got %f", tc.a, tc.b, tc.want, got)

		}

	}

}




func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 4, b: 2, want: 2},
		{a: 9, b: 4, want: 5},
		{a: 10, b: 6, want: 4},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f,%f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}


func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 8, b:2, want: 16},
		{a: 6, b:2, want: 12},
		{a: 3, b:3, want: 9},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f,%f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b:2, want: 1, errExpected: true},
		{a: 9, b:3, want: 3, errExpected: true},
		{a: 12, b:4, want: 3, errExpected: true},
	}
	for _, tc := range testCases {
		err :=
		errReceived := err != nil
		got, _ := calculator.Divide(tc.a, tc.b )
		//if tc.want != got {
		// t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		if tc.errExpected != (err != nil) {
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, errReceived)
		}
	}
}