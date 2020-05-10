# 10 - Planning to Fail

![](../img/mistake.svg)

Now that you're happy writing multiple tests with multiple test cases, let's bring together the work you've done in the last few exercises.

**TASK:** Create a `calculator_test.go` file containing all four of the calculator tests you've written so far:

* `TestAdd`
* `TestSubtract`
* `TestMultiply`
* `TestDivide`

You don't need to write anything new here; just copy and paste the code from the previous two exercises. Do the same for the `calculator` package, so that we have a calculator with `Add`, `Subtract`, `Multiply`, and `Divide` functions, all tested. (Don't forget to create a `go.mod` file!)

Recall that in the previous exercise we had to avoid any test inputs to `Divide` where `b` is zero. This time, we're going to extend `TestDivide` so that it deliberately triggers a division by zero, and checks that the `Divide` function handles this correctly.

So what _should_ the `Divide` function do in this case? Not crash the program, that's for sure. Instead, we'll make it return an error. In fact, `Divide` will return _two_ values. Right now it looks something like this:

```go
func Divide(a, b int) int {
	// This is where the magic happens!
}
```

We'll be updating the function so that it looks like this:

```go
func Divide(a, b int) (int, error) {
```

There are now two possibilities:

1. If `b` is zero, we can't do the division, so we return zero for the result, plus an error value containing a suitable message explaining what's wrong (you can use `errors.New(...)` to create this error value).

2. On the other hand, if `b` is non-zero, then we can do the division as normal, and return the result, and the accompanying error value will be `nil`.

You don't need to do anything to the `Divide` function yet, though. When we're developing test-first, we always start by updating the test to test the new functionality that we want. Only then do we actually implement that functionality.

So how _do_ we test for this? Well, in just the same way as we've done before: by calling the function, and comparing the results against what's expected. But this time, we have _two_ results, and _two_ expectations.

A simple way to do this would be to extend the test case struct like this:

```go
struct {
	a, b int
	want int
	err error
}
```

We could call the function, and check that the first value returned is equal to `want`, and the second value equal to `err`. But that's not ideal, because we'd have to write a test case like:

```go
{a: 1, b: 0, want: 0, err: errors.New("division by zero")},
```

What's wrong with that? Well, if we wanted to update the code at some point to return a different error message, then we'd have to update all the test cases as well. And there's really no need to test the _specific_ error message returned, is there? All we actually want to test is that there _is_ an error; that is, that the error result is not `nil`.

So a better test case structure is:

```go
struct {
	a, b int
	want int
	errExpected bool
}
```

For each test case, we're just specifying that an error is expected, or not. For example:

```go
{a: 1, b: 0, want: 0, errExpected: true},
```

**TASK:** Update the existing test cases for `TestDivide` to add the new `errExpected` field. We don't have any divide-by-zero cases at the moment, so `errExpected` should be `false` for every case.

## Errors and expectations

How do we test whether the function is returning errors when it's supposed to (and not returning them when it's not supposed to)? The first thing we need to do is to update the code to account for the fact that `Divide` now returns two values: the _data_ value (the answer to the sum, if there is one), and the _error_ value (indicating whether or not there was an error)

There's a convention in Go that, when a function returns both a data and an error value, we should ignore the data value if there's an error. In other words, the presence of an error means that the data value is invalid and shouldn't be used. That makes sense: if we try to divide a number by zero, for example, not only is that an error, but there's no meaningful data value we could return.

It's clear from this that when we're testing functions that can return errors, we need to _check the error value first_. If there's an error, we _need not_ check the data value against expectations, and in fact we _must not_, because it doesn't even make sense to expect anything in this case.

So there are four possible outcomes for a given test case:

1. The error value is `nil` (indicating no error), and the data value is what we expected (pass)
2. The error is `nil`, but the data value is wrong (fail)
3. The error is not `nil`, and `errExpected` is true, indicating that we expected an error (pass)
4. The error is not `nil`, but `errExpected` is false (fail)

**TASK:** Update the code of `TestDivide` so that it receives and checks the error value that will be returned from `Divide`, and fails the test if we _expected_ an error but didn't get one, or we _got_ an error we didn't expect.
Provided the error is non-nil, the test should also check the data value.

Don't try running the tests yet, because they won't work: we didn't yet update `Divide` to return the new error value we've added to our test cases!

**TASK:** Update `Divide` to return a suitable error value along with the data value. In the case of an error, as we've seen, it doesn't matter what data value we actually return, but the convention is to return zero. For example:

```go
return 0, errors.New("can't divide by zero")
```

The tests should now pass!

**TASK:** Check your test logic by changing one of the test cases to expect an error (set `errExpected: true`). This should cause the test case to fail. If not, there's something wrong!

Return the test case to its original state once you've checked this.

**TASK:** Add a new test case that triggers a divide-by-zero error, and make _that_ pass!

Check that it really works by setting the `errExpected` field on the test case to `false`. If the test still passes even when there's an unexpected error, there's something wrong with the test!

Once you're satisfied that the test fails correctly when it should, fix the test data again so that the test passes.

## Done?

When your tests pass, you're done! That's it for this set of exercises.
Nice work!

It's time to take a break. Why not reward yourself with a piece of cake, or a tasty snack of your choice?

If you'd like some further reading, try:

* [Go tutorials from Bitfield Consulting](https://bitfieldconsulting.com/golang)