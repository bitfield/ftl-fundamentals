# 03 - Fixing a Hole

![](../img/chained.png)

You've been called in to help another Gopher on a project that has a failing test. You'll need to run the tests, figure out what's wrong, fix the problem, and make sure the tests pass.

**TASK:** Run the command:

**`go test`**

```
--- FAIL: TestHello (0.00s)
    hello_test.go:12: want "Hello world!", got "Hello Gophers!"
FAIL
exit status 1
FAIL    hello   0.113s
```

This test failure is telling us exactly where the problem occurred:

```
hello_test.go:12
```

It also tells us what the problem was:

```
want "Hello world!", got "Hello Gophers!"
```

**TASK:** Open up the [`hello_test.go`](hello_test.go) file, and find the failing test:

```go
func TestHello(t *testing.T) {
	want := "Hello world!"
	got := hello.Greeting()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
```

This is a test of the [`hello.Greeting`](hello.go#L5) function. First, the test sets up a variable `want` to establish what it _wants_ to receive from the function:

```go
want := "Hello world!"
```

Then it calls the function and puts what it actually _got_ into another variable:

```go
got := hello.Greeting()
```

Then it checks to see if `got` is the same as `want`. If the two values are different, then the `hello.Greeting` function must have returned something unexpected:

```go
if want != got {
```

And if that's the case, the test fails with a message that tells us what it wanted and what it got:

```go
t.Errorf("want %q, got %q", want, got)
```

On the other hand, if `want` and `got` are equal, then everything's fine, so the test function ends here.

The fact that we are seeing the test failure message means that `hello.Greeting` returned something other than the expected `"Hello world!"`. In fact, it's returning `"Hello, Gophers!"`. That's close, but not close enough.

**TASK:** Open up the [`hello.go`](hello.go) file and find the `Greeting` function. See if you can modify it to return the expected result.

Run `go test` again to check that you got it right.

## Done?

When the test passes, you're done! Nice work. Go on to the next exercise:

* [Say Hello to my Little Friend](../04/README.md)

## What's this?

This is one of a set of introductory Go exercises by [John Arundel](https://bitfieldconsulting.com/golang/learn) called [For the Love of Go - Fundamentals](../README.md).

<small>Gopher image by [egonelbre](https://github.com/egonelbre/gophers)</small>