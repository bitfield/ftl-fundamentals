In previous exercises, we've tested just one function at a time, even when we tested it with multiple cases. This time, we're going to write _three_ tests, one for `calculator.Subtract`, one for `calculator.Multiply`, and another for `calculator.Divide`.

To get you started, the `calculator_test.go` file contains three _stub_ tests that you'll need to complete:

```go
func TestSubtract(t *testing.T) {
	t.Fatal("Not implemented yet")
}

func TestMultiply(t *testing.T) {
	t.Fatal("Not implemented yet")
}

func TestDivide(t *testing.T) {
	t.Fatal("Not implemented yet")
}
```

Turn these into table-driven tests, with ten test cases each, just like in the previous exercise, and make them pass.

For TestDivide, don't include any test cases that require dividing by zero. If we try to divide a number by zero in Go, we get a runtime error:

```
division by zero
```

We'll see how to handle this in the next exercise.