package calculator_test

import (
	"calculator"
	"testing"
)

func TestCalculator(t *testing.T) {

	var tests = []struct {
		function func(float64, float64) float64
		want     float64
		input    []float64
	}{
		{function: calculator.Add, input: []float64{2, 2}, want: 4},
		{function: calculator.Subtract, input: []float64{2, 4}, want: 2},
		{function: calculator.Multiply, input: []float64{2, 4}, want: 8},
	}
	for _, tt := range tests {
		got := tt.function(tt.input[0], tt.input[1])
		if tt.want != got {
			t.Errorf("want %f, got %f", tt.want, got)
		}

	}
}

func TestDivide(t *testing.T) {
	var tests = []struct {
		a             float64
		b             float64
		want          float64
		errorExpected bool
	}{
		{
			a:             4,
			b:             0,
			want:          999,
			errorExpected: true},
		{
			a:             4,
			b:             2,
			want:          2,
			errorExpected: false,
		},
	}
	for _, tt := range tests {
		got, err := calculator.Divide(tt.a, tt.b)
		if tt.errorExpected && err == nil {
			t.Fatal("want err/0, got nil")
		}
		if !tt.errorExpected && err != nil {
			t.Fatal("Unexpected error")
		}
		if !tt.errorExpected && tt.want != got {
			t.Errorf("given %f/%f, want: %f, got: %f", tt.a, tt.b, tt.want, got)
		}

	}

}

func TestSquare(t *testing.T) {
	var tests = []struct {
		a             float64
		want          float64
		errorExpected bool
	}{
		{
			a:    2,
			want: 4,
		},
		{
			a:    -2,
			want: 4,
		},
	}
	for _, tt := range tests {
		got := calculator.Square(tt.a)
		if tt.want != got {
			t.Errorf("given %f, want: %f, got: %f", tt.a, tt.want, got)
		}

	}

}
