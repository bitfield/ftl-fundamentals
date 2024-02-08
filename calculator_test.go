package calculator_test

import (
	"calculator"
	"fmt"
	"math/big"
	"testing"
)

const (
	SHLD_SUCCEED = 1 // We expect the test to succeed. If it fails, its an error
	SHLD_FAIL    = 0 // We expect the test to fail. It it succeeds, its an error
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
		{2, 2.030000, 4.030000, "add fractions test 1: exp_succ", SHLD_SUCCEED},
		{6, 3.99, 9.99, "add fractions test 2: exp_succ", SHLD_SUCCEED},
		{2, 2.03675, 4.03675, "add fractions test3: exp_succ", SHLD_SUCCEED},
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
		{2.030000, 20, 9.852216, "Divide fractions 1", SHLD_SUCCEED},
		{2, 0, 0, "Divide fractions 3", SHLD_SUCCEED},
		{0, 3.99, 0, "Divide by 0", SHLD_FAIL},
	}
	cntTestExec, cntfuncErr, cntexpMismatch := 0, 0, 0
	cntNoOfTests := len(tCases)
	// If I have deferred it, why is it running here, in the flow. How is it related to tests?
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
		case SHLD_FAIL: // We have defined this test case to FAIL. Its err if it doesn't
			if !(funcErr || expMismatch) {
				t.Errorf("DivTest: %s :: want %f, got %f ", tc.tcName, tc.want, got)
			}
		case SHLD_SUCCEED: // We have defined this test case to SUCCEED. Its err if it doesn't
			if funcErr || expMismatch {
				t.Errorf("DivTest: %s :: want %f, got %f ", tc.tcName, tc.want, got)
			}
		}
	}
	fmt.Printf("SummaryDiv: TotalTests: %d: Executed: %d: Failed: %d: Succeeded: %d\n", cntNoOfTests, cntTestExec, cntfuncErr, cntexpMismatch)

}

func printDivTestSummary(cntNoOfTests, cntTestExec, cntfuncErr, cntexpMismatch int) {
	fmt.Printf("SummaryDiv: TotalTests: %d: Executed: %d: Failed: %d: Succeeded: %d\n", cntNoOfTests, cntTestExec, cntfuncErr, cntexpMismatch)
}
