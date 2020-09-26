package calculator_test

import (
	"calculator"
	"fmt"
	"math/big"
	"testing"
)

const (
	SHLD_SUCCEED = 1
	SHLD_FAIL    = 0
)

type testCase struct {
	a           float64
	b           float64
	want        float64
	tcName      string
	tcExpStatus int
}

func TestAdd(t *testing.T) {
	t.Parallel()
	tCases := []testCase{
		{2, 2.030000, 4.030000, "some fractions1", SHLD_SUCCEED},
		{6, 3.99, 9.99, "some fractions2", SHLD_SUCCEED},
		{2, 2.03675, 4.03675, "some fractions3", SHLD_SUCCEED},
	}
	for _, tc := range tCases {
		got := calculator.Add(tc.a, tc.b)
		var bgot = big.NewFloat(got)
		var bwant = big.NewFloat(tc.want)
		if result := bgot.Cmp(bwant); result != 0 {
			t.Errorf("Add: Test: %s : want %f, got %f and Cmp gave %d ", tc.tcName, tc.want, got, result)
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
		{2.030000, 20, 9.852216, "Divide some fractions1", SHLD_SUCCEED},
		{0, 3.99, 0, "Divide some fractions2", SHLD_FAIL},
		{2, 0, 0, "Divide some fractions3", SHLD_SUCCEED},
	}
	cntTestExec, cntfuncErr, cntexpMismatch := 0, 0, 0
	cntNoOfTests := len(tCases)
	defer printDivTestSummary(cntNoOfTests, cntTestExec, cntfuncErr, cntexpMismatch)

	for _, tc := range tCases {
		funcErr := false
		expMismatch := false
		cntTestExec++

		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			funcErr = true
			cntfuncErr++
			t.Fatalf("Test: %s :: (%f / %f) : Divide by Zero", tc.tcName, tc.b, tc.a)
		} else {
			if tc.want != got {
				expMismatch = true
				cntexpMismatch++
			}
		}
		switch tc.tcExpStatus {
		case SHLD_FAIL: // We have defined this test case to FAIL. Its err it it doesn't
			if !(funcErr || expMismatch) {
				t.Errorf("want %f, got %f :: Test: %s", tc.want, got, tc.tcName)
			}
		case SHLD_SUCCEED: // We have defined this test case to SUCCEED. Its err it it doesn't
			if funcErr || expMismatch {
				t.Errorf("want %f, got %f :: Test: %s", tc.want, got, tc.tcName)
			}
		}
	}
	fmt.Printf("SummaryDiv: TotalTests: %d: Executed: %d: Failed: %d: Succeeded: %d\n", cntNoOfTests, cntTestExec, cntfuncErr, cntexpMismatch)

}

func printDivTestSummary(cntNoOfTests, cntTestExec, cntfuncErr, cntexpMismatch) {
	fmt.Printf("SummaryDiv: TotalTests: %d: Executed: %d: Failed: %d: Succeeded: %d\n", cntNoOfTests, cntTestExec, cntfuncErr, cntexpMismatch)
}
