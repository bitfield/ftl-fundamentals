# Fundamentals of Go

![](img/witch-learning.png)

Hello, and welcome to learning Go! It's great to have you here.

## What's this?

This is a set of introductory exercises in Go programming, suitable for complete beginners. If you don't know anything about Go yet, or programming, but would like to learn, you're in the right place!

(If you do already know something about Go, you should find these exercises relatively straightforward, but still worth doing.)

## What you'll need

You'll need to install Go on your computer, if you don't have it already (if you can run the command `go version` without errors, you're all set!)

Follow the instructions here to download and install Go:

* [https://golang.org/dl/](https://golang.org/dl/)

While all you need to write and run Go programs is a terminal and a text editor, you'll find it very helpful to use an editor which has specific support for Go. For example, [Visual Studio Code](https://code.visualstudio.com/) has excellent Go integration.

## What you'll learn

In this set of exercises, you'll learn:

* How to run tests for a Go program
* How to automatically format your Go code correctly
* The basic pattern that all Go tests should follow
* How to declare and import Go _packages_ (units of code)
* How to design multiple _test cases_ and use them in your tests
* How to test error handling in your programs
* A simple, reliable, test-first development workflow

## Test-driven development

There's a style of programming (which isn't unique to Go, but is very popular with Go developers) called _Test-Driven Development_ (TDD).

What this means is that before we write a program to do something (multiply two numbers, let's say), we first of all write a _test_.

A test is also a program, but it's a program specifically designed to run _another_ program with various different inputs, and check its result. The test verifies that the program actually behaves the way it's supposed to.

For example, if we were considering writing a function called `Multiply`, and we wanted to be able to check that it worked, we might write a program that calls `Multiply(2, 2)` and checks that the answer is `4`. That's a test!

Lots of programmers use tests, but without necessarily doing TDD. What makes TDD, in particular, so useful, is that we actually write the tests _first_. In most of these exercises, you'll start by writing a test for a particular program, then you'll write the program itself. That might seem backwards to you, but it makes a lot of sense!

By writing the test first, we are forced to think about what the program actually needs to do. We also have to design its interface (the way users interact with the program), since the test interacts with the program the same way a user would. And it also means we know when to stop coding, because when the test starts passing, we're done!

(You might ask why not write the code first, and then, once the code is working, write the test for it. Well, how would you know whether or not it's working without a test? You see, the whole scheme collapses. Write your tests first.)

## How to do the exercises

Each exercise is described in a separate section of this README file. Each section briefly explains what the exercise is about, and outlines what you'll be doing. There'll be one or more goals for you to achieve, marked with **GOAL:** in bold. (Don't worry if you're not sure how to do it. There'll usually be some helpful suggestions following the goal description.)

If you're finding the goals easy, or want a bit more of a challenge, there are some _stretch goals_, too. These are optional; if you're not interested or not ready to tackle them, just skip over them.

So let's get started!

# 1: Testing times

![](img/calculator.jpg)

It's your first day as a Go developer at Texio Instronics, a major manufacturer of widgets and what-nots, and you've been assigned to the Future Projects division to work on a very exciting new product: a Go-powered electronic calculator! Welcome aboard.

In this exercise, all you need to do is make sure your Go environment is set up and everything is working right.

One of your new colleagues has already made a start on the calculator project.  She's placed a Go package in the [`calculator.go`](calculator.go) file, and an accompanying test in [`calculator_test.go`](calculator_test.go). You're going to run this test and make sure it passes.

**GOAL:** Successfully run the test.

Here's how to do that:

In this directory, run the command:

**`go test`**

If everything is good, we will see this output:

```
PASS
ok      calculator      0.234s
```

This tells you that all the tests passed in the package called `calculator`, and that running them took 0.186 seconds. (It might take a longer or shorter time on your computer, but that's okay.)

If you see test failure messages, or any other error messages, there may be a problem with your Go setup. Don't worry, everybody runs into this kind of issue at first. Try Googling the error message to see if you can figure out what the problem is.

Once the test is passing, you can move on!

**STRETCH GOAL:** Use your editor's Go support features to run the test directly in the editor. For example, in Visual Studio Code, there's a little 'Run test' link above each test function, and a 'Run package tests' link at the top of the file. Find out how to do this in the editor you're using.

# 2: Getting in shape

![](img/hiking.png)

The next thing to do is ensure that your Go code is formatted correctly. All Go code is formatted in a standard way, using a tool called `gofmt` (pronounced, in case you were wondering, 'go-fumpt'). In this exercise you'll see how to use it to check and reformat your code (if necessary).

The file `calculator.go` in this directory contains a deliberate formatting mistake, just to make sure you're on your toes!

**GOAL:** Run `gofmt` to check the formatting and find out what is wrong:

**`gofmt -d calculator.go`**

```
diff -u calculator.go.orig calculator.go
--- calculator.go.orig  2020-08-11 15:13:42.000000000 +0100
+++ calculator.go       2020-08-11 15:13:42.000000000 +0100
@@ -3,5 +3,5 @@

 // Add takes two numbers and returns the result of adding them together.
 func Add(a, b float64) float64 {
-return a + b
+       return a + b
 }
```

Look at the lines beginning with `-` and `+`:

```
-return a + b
+       return a + b
```

`gofmt` is telling us what it wants to do. It wants to remove (`-`) the line that has no indent before the `return` statement, and replace it (`+`) with a line where the statement is properly indented (by one tab character).

We could make this change using our editor, but there's no need: `gofmt` can do it for us.

**GOAL:** Run:

**`gofmt -w calculator.go`**

This will reformat the file in place. You can check that it's worked by running `gofmt -d` on the file again. Now there should be no output, because it's already formatted correctly.

If you're using an editor that has Go support, such as Vim, Visual Studio Code, or GoLand, you can set up the editor to automatically run your code through `gofmt` every time you save it. This will help you avoid forgetting to format your code. Check the documentation for your editor to see how to set this up.

**STRETCH GOAL:** Experiment with `gofmt`. Try formatting your code in different ways and see what changes `gofmt` makes to it. It's interesting to note what it does and doesn't change.

# 3: Fixing a hole

![](img/chained.png)

Now that your development environment is all set up, your colleague needs your help. She has been working on getting the calculator to subtract numbers, but there's a problem: the test is not passing. Can you help?

You'll find the test code commented out in the `calculator_test.go` file:

```go
// func TestSubtract(t *testing.T) {
// 	t.Parallel()
// 	var want float64 = 2
// 	got := calculator.Subtract(4, 2)
// 	if want != got {
// 		t.Errorf("want %f, got %f", want, got)
// 	}
// }

```

**GOAL:** Get this test passing!

Uncomment the test function and run the tests:

**`go test`**

```
--- FAIL: TestSubtract (0.00s)
    calculator_test.go:20: want 2.000000, got 6.000000
FAIL
exit status 1
FAIL    calculator      0.178s
```

This test failure is telling us exactly where the problem occurred:

```
calculator_test.go:20
```

It also tells us what the problem was:

```
want 2.000000, got 6.000000
```

So what exactly is the test doing? Let's take a closer look. Skip the `t.Parallel()` part for now, as it merely relates to how Go runs the tests.

First, the test sets up a variable `want` to establish what it _wants_ to receive from the function:

```go
	var want float64 = 2
```

Then it calls the function with the values `4, 2`, and stores the result into another variable `got`:

```go
	got := calculator.Subtract(4, 2)
```

The idea is to compare `want` with `got` and see if they are different. If they are, the `Subtract()` function is not working as expected, so the test fails:

```go
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
```

On the other hand, if `want` and `got` are equal, then everything's fine, so the test function ends here. (You don't need to explicitly make a test pass in Go; if it doesn't explicitly fail, that's considered a pass.)

So let's look at the code for the `Subtract()` function (in [`calculator.go`](calculator.go)). Here it is:

```go
// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return b - a
}
```

If you spot a problem, try altering the code to fix it. Run `go test` again to check that you got it right.

When the test passes, you can move on!

**STRETCH GOAL:** Find out what other ways there are to fail a test using the `testing.T`. What other methods does it have? Do any of them look potentially useful?

# 4: Go forth and multiply

![](img/rocket.png)

Excellent work! You now have a calculator which can add and (correctly) subtract. That's a great start. Let's turn to multiplication now.

Up to now you've been running existing tests and modifying existing code. For the first time you're going to write the test, and the function it's testing!

**GOAL:** Write a test for a function `Multiply()` which, just like the `Add()` and `Subtract()` functions, takes two numbers as parameters, and returns a single number representing the result.

Where should you start? Well, this is a test, so start in the [calculator_test.go](calculator_test.go) file. Test functions in Go have to have a name that starts with `Test` (or Go won't call them when you run `go test`). So `TestMultiply` would be a good name, wouldn't it?

You'll see that `TestAdd` and `TestSubtract` look much the same, except for the specific inputs and the expected return value. So start by copying one of those functions, renaming it `TestMultiply`, and making the appropriate changes.

When you're done, running the tests should produce a compilation error:

```
undefined: calculator.Multiply
```

This makes sense; you haven't written that function yet!

**GOAL:** Write the minimum code necessary to get the test to compile and fail.

What is the minimum code necessary to compile? Well, you need to _define_ the `Multiply` function, if not implement it. See if you can work out what its function signature should be, based on looking at `Add` and `Subtract`.

It will take two `float64` parameters which are its inputs, and since the function must return something if it's to be useful, it needs to be defined as returning a single `float64` value.

Given that function signature, the function will not compile unless it has a `return` statement, and the easiest way to write this is simply to have the function return zero. This should be enough to make the program compile, and if you run `go test` again, `TestMultiply` should fail.

## Why deliberately write incorrect code?

Why do this? Why write something in `Multiply` which is obviously incorrect; namely, returning zero? Why not go ahead and actually write the code to do the multiplication?

This is an important step in test-driven development. We start by writing a test for the function that doesn't yet exist. Before we start writing the function itself, we need to verify that the test is correct! (Otherwise, we might end up with _both_ the test _and_ the tested function being wrong.)

One way to do this is to ensure that, with a minimal and obviously incorrect implementation of `Multiply`, we see the test fail the expected way. If the test passes, even though we know `Multiply` returns zero for all inputs, something is really wrong with the test!

Once you have `TestMultiply` failing correctly (that sounds wrong, but you know what I mean), you can move on to the next step.

**GOAL:** Write the minimum code necessary to get `TestMultiply` to pass.

Why did we say "the minimum code necessary"? Because, in a sense, the test itself defines what `Multiply` needs to do. Once it does that, and the test passes, we should _stop coding_.

## Why stop when the test passes?

Why? Because any further code we write after the test passes, is by definition untested! We have no way of knowing whether it's correct, because the test passes whether the extra code is there or not. So we don't do that. We write code until the test passes, and then we stop.

**GOAL:** Make `TestMultiply` fail deliberately.

There's one final step required for us to be really sure that both the test and the code are correct. We need to see the test fail again. To do this, change the expected return value from `Multiply` to something that you know is wrong, and run `go test` again.

## Why break a passing test?

Why do this? Well, during the implementation of `Multiply` you may have tweaked the test code a little. It's possible to have a passing test which passes incorrectly! For example, if you accidentally wrote `if want == got` instead of `if want != got`.

If you've already proved that `Multiply` is correct, and then you change the test to expect an _incorrect_ result from `Multiply`, and it still passes, then, again, there's something really wrong with the test. (This step is also useful because you'll see the failure message that the test outputs. If it needs rewording or expanding with further information, this is a good time to do that.)

Once you're satisfied that the broken test fails correctly, restore it to the way it was, and check that it now _passes_ again.

**STRETCH GOAL:** Think about other ways this test could be wrong. Try some of them out. What happens? What could you do to catch such potential problems?

Nice work! Go on to the next section.

# 5: Testing to destruction

![](img/crash-dummy.png)

Now that you've successfully designed and implemented a new calculator feature, test-first, it's time to think a little more about testing, and how you can extend it. The `Add`, `Subtract`, and `Multiply` functions are so simple that it's fairly difficult (though not at all impossible) to get them wrong. If `TestAdd` passes for one set of inputs, for example, it's hard to imagine other inputs that would make it fail.

But this doesn't apply to more complicated functions, where it's quite likely that they could work correctly for some inputs, but not for others. In this situation it will be a good idea to test multiple _cases_ for each function: that is to say, different inputs with different expected results.

For example, the current version of `TestAdd` tests calling `Add(2, 2)` and checks that the result is 4. You could write a new function `TestAdd1and1`, for example, which calls `Add(1, 1)` and expects 2. But this would be tedious for many test cases.

## Introducing test cases

Instead, the most elegant solution would be to define a number of different _test cases_—each representing a pair of inputs, together with the corresponding expected output—and then loop over them, calling `Add` with the inputs and checking the output against that expected. If _any_ test case fails, the whole test fails.

Let's see what that might look like. First, it will be convenient to define a `struct` type to represent the test case:

```go
type testCase struct {
	a, b float64
	want float64
}
```

Here's the test case you already have, expressed in this new form:

```go
testCase{
	a: 2,
	b: 2,
	want: 4,
}
```

## A slice of test cases

In the test function, you could declare a slice of these test cases with something like the following (to make these more concise, let's write each `testCase` literal on a single line):

```go
testCases := []testCase{
	{ a: 2, b: 2, want: 4 },
	{ a: 1, b: 1, want: 2 },
	{ a: 5, b: 0, want: 5 },
}
```

Make sure you're confident you understand what's happening here before moving on. You're setting up a slice called `testCases`, with three elements. Each element of this slice is an instance of the `testCase` struct type, which has the fields `a`, `b`, and `want`, representing the inputs and expected outputs of the test case.

## Looping over test cases

Now that you have the test cases, you can write the code to loop over them. That's straightforward; you can use the `range` operator to do this.

```go
for _, tc := range testCases {
	... // call Add() with the inputs and check the result
}
```

You already know how to do the part inside the loop, because it's just the same as the existing `TestAdd` function, only with altered variable names:

```go
	got := calculator.Add(tc.a, tc.b)
	if tc.want != got {
		t.Errorf("want %f, got %f", tc.want, got)
	}
```

Each time through this code, `tc` will be the next `testCase` struct, so `tc.a` and `tc.b` will be the inputs to `Add`, and `tc.want` will be the expected result. Let's also extend the failure message to include the inputs, since these will be different for each test case:

```go
t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
```

What would that look like? Suppose the `1 + 1 == 2` case failed. You would see this message:

```
Add(1.000000, 1.000000): want 2.000000, got 3.000000
```

Very helpful!

**GOAL:** Rewrite `TestAdd`, `TestSubtract`, and `TestMultiply` to test multiple cases for each function. Try to think of interesting and varied test cases: for example, negative numbers, zeroes, fractional numbers.

As before, once each test is passing, break it deliberately and make sure it fails in the expected way, with a helpful failure message.

**STRETCH GOAL:** Add a string field `name` to your test cases, and supply a human-readable name for each case. For example, a test case for `Add` might be named `"Two negative numbers which sum to a positive"`. Output the name of the test case along with the failure message.

When you're happy with the expanded tests, move on to the next section.

# 6: Errors and expectations

![](img/mistake.png)

The board reviewed the calculator project yesterday, and they're really pleased with your progress! The only remaining major feature to be completed is division, and that needs a little more thought than the others.

Why? Because it's possible to write division expressions that have _undefined_ results. For example:

```
6 / 0
```

This looks reasonable at first sight, but what we're really asking here is "What number, multiplied by zero, gives 6?" And clearly there is no such number. What should our calculator do in this case?

## Error values in Go

There are several options: it could give the answer `0`, which seems unsatisfying, or it could give an answer that represents infinity (perhaps the symbol ∞), or it could define a symbol that represents 'not a number' (the abbreviation `NaN` is sometimes used for this.)

But these are awkward attempts to work around a fundamental issue: dividing by zero is _not allowed_ in the ordinary rules of arithmetic. So there's something about the `Divide` function which wasn't true of `Add`, `Subtract`, and `Multiply`: it's possible to give it _invalid input_.

A cool thing about functions in Go is that they can return multiple values (and often do). So we could have `Divide` return two values: one can be the result of the calculation, as with the other functions, but the other can be an indicator saying whether or not the input was valid.

## Functions that return errors

The standard way to do this in Go is to write a function signature like:

```go
func Divide(a, b float64) (float64, error)
```

It takes two inputs `a` and `b` of type `float64`, as with the other arithmetic functions, but unlike them, it returns two values: a `float64` (the answer) and a value of type `error` (the 'invalid input' indicator).

So how would this work? In the case of well-formed divisions (3 divided by 2, let's say), then `Divide` would return `1.5` as the answer (let's call this the 'data value') and `nil` as the error indicator (let's call this the 'error value'). `nil` means 'no error'—everything's fine.

Now let's take the case of dividing by zero. Instead of returning `nil` for the error value, we would return something else. (It doesn't really matter what; literally anything other than `nil` sends the message that there was an error.)

## What to return when there's an error

What would we return for the data value? Well, that doesn't matter either, and there is no right answer. Conventionally, we would just return zero in this case, but that's arbitrary: the presence of a non-nil error value is a warning that says "Ignore the data value. It's useless, because there was an error!"

**GOAL:** Define a test case struct type which allows us to test both valid and invalid inputs to `Divide`.

How would we need to modify the struct type we've been using up to now? Well, for functions with two inputs and one output, we had a type like this:

```go
type testCase struct {
	a, b float64
	want float64
}
```

Two inputs, and one expected output. Now we have a function with two inputs and _two_ outputs. Can you see what we need to add to the struct?

We need a way for the test case to specify whether or not the given inputs should produce an error. The simplest way to do that is to use a `bool` value, which we could call something like `errExpected`. If this value is `true`, it's saying the test case should cause `Divide` to return a non-nil error. For example, here's our division-by-zero test case expressed in this form:

```go
{ a: 6, b: 0, want: 0, errExpected: true }
```

**GOAL:** Write a set of test cases for `Divide` using the new struct type. Make sure to include both valid and invalid inputs.

## Testing for error status

Now that we have a set of suitable test cases, let's think about how to write the `TestDivide` function itself. Previously, there were only two possible situations we had to deal with:

1. The data value didn't match the value we expected (fail)
2. The data value matched the expectation (pass)

Now there are _four_ possible outcomes:

1. We expected an error, but we didn't get one (fail)
2. We _didn't_ expect an error, but _did_ get one (fail)
3. The error value was as we expected, but the data value didn't match the expectation (fail)
4. The error value was as we expected, and the data value also matched expectation (pass)

It might look a little complicated, but it's really quite straightforward, if you've followed everything so far. Make sure you're happy with these four outcomes before moving on.

**GOAL:** Write the necessary logic for the `TestDivide` function, handling each of the four possible outcomes correctly.

## Checking for unexpected error status

This needs some care, because it's very easy to get the error-checking logic wrong. Consider this situation, for example (Outcome 1): we expected an error, but we got none (we got a `nil` error value). What should the failure message be for this outcome?

```
Divide(6, 0): unexpected error status: nil
```

This would be a reasonable output. What we _don't_ want to see is something like this:

```
Divide(6, 0): want 0, got ...
```

Why not? Because the reason the test is failing is that the function is _not returning an error_ when it's supposed to. In that situation, the data value is entirely irrelevant, and comparing it against any expectation would be a waste of time. Worse, it would give a misleading message, making it harder to debug the problem. (A good way to catch this kind of issue is to give _obviously crazy_ `want` values for your error test cases, such as `999`. If you see these values in your test outputs, something's wrong with the test, because in the case of unexpected error status, it shouldn't even check the data value.)

The implication of this is that the test needs to check the error value against expectation _before even looking at the data value_. Here's one way to do that:

```go
errReceived := err != nil
if tc.errExpected != errReceived {
	t.Fatalf("Divide(%f, %f): unexpected error status: %v", errReceived)
}
```

This neatly deals with both outcomes 1 and 2, which both involve "unexpected error status"; either because we expected one but didn't get one, or didn't expect one and did get one!

What is `errReceived`? It's a boolean value which tells us whether or not we got an error from `Divide`. If you think about it, the value of this variable needs to be the same as the test case's `errExpected` field. If not, it's a test failure.

A shorter way to write this logic is:

```go
if tc.errExpected != (err != nil) {
	t.Fatalf(...)
}
```

and you'll sometimes see this in Go code. However, this isn't as clear as explicitly setting `errReceived`, and I recommend you be as explicit as possible, especially in tests.

## Failing and bailing

Why call `t.Fatalf()` in this case, rather than `t.Errorf()`? What's the difference?

Calling `t.Errorf()` outputs the failure message, but the test function will continue executing (it might go on to make other comparisons, for example). On the other hand, `t.Fatalf()` outputs its failure message and exits the test function immediately. It fails and bails, you might say.

Bailing is useful when we have established that things are so broken that there's no point continuing with the test. For example, if the error status from the function is unexpected, then comparing the data value is pointless (indeed, it would be misleading about the causes of the failure). So `t.Fatalf()` allows us to skip any further checks within this test (though it doesn't stop _other_ test functions from being run. It bails out of _this_ test, not all tests.)

Another common use of `t.Fatalf()` is when there's some error setting up the test itself. For example, suppose we want to compare the output of a function with the contents of some file (known as a _golden file_, and conventionally placed in a `testdata/` directory). We might start by reading this file from disk, in preparation for comparing it with the test output. But if reading the file fails for some reason (file missing, bad permissions, out of memory; life is full of unexpected snags) then clearly there's no point continuing with this test. Bailing out with `t.Fatalf()` (and some diagnostic information) is appropriate here.

## Checking the data value

There's one thing left to do in this test, and it's something we've done before: checking the data value returned from the function against our expectation.

However, now there's an extra subtlety, because we're dealing with a function that can error. When we're testing a case which is _supposed_ to error, we don't want to check the data value. No expectation for it would be meaningful (indeed, if we're wily enough, we have set an obviously wrong expectation of `999`, or something like that, to detect precisely this problem).

Therefore, once we have established that the error status is as expected (if not, we've already failed and bailed), we can write:

```go
if !tc.errExpected && tc.want != got {
```

In other words, compare the data value _only if the test case does not expect an error_.

**GOAL:** Write the minimum code necessary to make `TestDivide` pass.

When your test cases all pass, and you've verified them all by deliberately breaking their expectations, and you've checked that the test outputs the correct failure message for unexpected error statuses, then you can move on to the next (and final) section.

**STRETCH GOAL:** Write a new test for one or more of your functions which generates _random_ test inputs instead of using prepared cases (you can use the `math/rand` library for this). For example, you might write a `TestAddRandom` function which generates two random `float64` values, sums them to get the expected result (_not_ using the `Add` function!), and calls `Add` to verify that it returns the same answer. Do this many times, say, a hundred times, to give the function a good workout.

# 7: Roots music

![](img/music.png)

Well done! That was a good piece of work getting the test cases and error handling set up for the `Divide` function. In fact, you now know just about everything you need to write great tests in Go!

## The structure of tests

All tests follow more or less the same basic structure:

* Define the test cases with their inputs and expected outputs (including any error outputs)
* Iterate over the test cases, calling the function under test with each set of inputs
* Check that the outputs match expectations, and fail otherwise with some helpful output

And you've also practiced a specific workflow, or sequence of actions, that you can use when developing any piece of Go code:

## The development process

1. Start by writing a test for the function (even though it doesn't exist yet)
2. See the test fail with a compilation error because the function doesn't exist yet.
3. Write the minimum code necessary to make the test compile and fail.
4. Make sure that the test fails for the reason you expected.
5. Write the minimum code necessary to make the test pass.
6. Change the test expectations to deliberately make the test fail, and check that the failure message is accurate and informative.
7. Restore the original expectations and make sure the test still passes.
8. Commit!

In practice, we may go through this loop a few times with any particular test. We start with something very simple, so that we are never facing an insurmountable problem. If we are, then we try to break that problem down into smaller problems, and repeat this process with those.

## A `Sqrt` function

Our basic calculator is complete, but you've just received an urgent message from the VP of Sales. She wants an extra premium feature which can be used to upsell users to the 'Enterprise' calculator. The ordinary calculator has been re-branded as the 'Home' edition.

It seems that enterprise users would like the ability to take square roots, so the VP's asking you to develop a `Sqrt` function. You already know everything you need to know to design, test, and develop this feature, so let's get started!

**GOAL:** Develop a `Sqrt` function which takes a `float64` value and returns its square root as a `float64` value.

You can use any standard or third-party libraries you think are appropriate. The function should return an error for negative inputs (no real number squared equals a negative number, so taking the square root of a negative number is not a valid operation). Your test should check the error handling in the same way as the test for `Divide`.

When the feature is complete to your own satisfaction, you're done! Nice work. The company has awarded you a sizeable bonus, and an all-expenses-paid vacation. Enjoy it!

**STRETCH GOAL:** Extend your existing tests and functions to allow _two or more_ inputs to `Add`, `Subtract`, `Multiply`, and `Divide`. For example, calling `Divide(12, 4, 3)` should divide 12 by 4, divide the result of that by 3, then return _that_ result. Calling `Multiply` with ten inputs should return the result of multiplying them all together.

**STRETCH GOAL 2:** Write a function which accepts strings of the form `12 * 3`, for example, and returns the result. You only need to handle expressions which have a single floating-point value, followed by a single operator, followed by a single floating-point value, ignoring any whitespace. For example, the following should all be valid inputs to your function:

```
2*2
1 + 1.5
18   /   6
100-0.1
```

## Who wrote this?

[John Arundel](https://bitfieldconsulting.com/about) is a Go teacher and mentor of many years experience. He's helped literally thousands of people to learn Go, with friendly, supportive, professional mentoring, and he can help you too! Find out more:

* [Learn Go with mentoring](https://bitfieldconsulting.com/golang/learn)

## What's next?

You might like to try another set of exercises in this series:

* [The G-machine](https://github.com/bitfield/gmachine) guides you through the development of a simplified virtual CPU in Go, complete with its own machine language. Along the way you'll learn some useful computer science fundamentals.

## Further reading

You can find more Go tutorials and exercises here:

* [Go tutorials from Bitfield](https://bitfieldconsulting.com/golang)

<small>Gopher images by the magnificent [egonelbre](https://github.com/egonelbre/gophers)</small>
