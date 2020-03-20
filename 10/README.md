Now that you're happy writing multiple tests with multiple test cases, let's bring together the work you've done in the last couple of exercises.

Create a `calculator_test.go` file containing all four of the calculator tests you've written so far:

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

Go ahead and make this change, and update the existing test cases for `TestDivide` to fit the new structure. We don't have any divide-by-zero cases at the moment, so `errExpected` should be `false` for every case. Make sure that your test checks both that the value returned by `Divide` matches `testCase.want`, and that the error value is `nil` if `testCase.errExpected` is false.

Naturally enough, the test will now fail, because we didn't yet update `Divide` to return the error value along with the result. Go ahead and do that now, so that the test passes.

Finally, add a divide-by-zero test case, and make _that_ pass! Check that it really works by setting the `errExpected` field to `false`. The test should fail now, because an error was not expected, but one was returned. If the test still passes even when there's an unexpected error, there's something wrong with the test!

Once you're satisfied that the test fails correctly when it should, fix the test data again so that the test passes.

Nice work. It's time to take a break. Why not reward yourself with a piece of cake, or a tasty snack of your choice?