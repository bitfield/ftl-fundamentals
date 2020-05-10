# 09 - Special Operations

![](../img/knight.svg)

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

**TASK:** Turn these into table-driven tests, with five test cases each, just like in the previous exercise, and make them pass. There are some special instructions for `TestDivide`, given below.

## TestDivide: Handle With Care

For TestDivide, don't include any test cases that require dividing by zero. If we try to divide a number by zero in Go, we get a runtime error:

```
division by zero
```

We'll see how to handle this in the next exercise.

Because our calculator only handles _integer_ arithmetic (that is, dealing with whole numbers), we'll also need to make sure that we don't include any divisions that result in a _fractional_ part. For example, the result of dividing 2 by 3 can't be expressed as an integer. Stick to test cases that have an integer result, for now.

## Done?

When all your test cases pass, you're done! Go on to the next exercise:

* [Planning to Fail](../10/README.md)

## What's this?

This is one of a set of introductory Go exercises by [John Arundel](https://bitfieldconsulting.com/golang/learn) called [For the Love of Go - Fundamentals](../README.md).

<small>Gopher image by [egonelbre](https://github.com/egonelbre/gophers)</small>