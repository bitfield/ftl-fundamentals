# 04 - Say Hello to my Little Friend

![](../img/zorro.svg)

We have another failing test. Your job, once again, is to make it pass!

**TASK:** Run the tests:

```
go test
# hello_test [hello.test]
./hello_test.go:10:9: undefined: hello.Greeting
FAIL    hello [build failed]
```

This time, there is no `hello.Greeting` function at all. You will need to create one, and have it return the result that the test expects. If you run into any trouble, have a look at the Greeting function in previous exercises and copy what it does, making suitable alterations to pass the test.

**TASK:** Add a `Greeting` function to the `hello` package, based on what the test expects.

## Done?

When the test passes, you're done! Go on to the next exercise:

* [Epic Fail](../05/README.md)

## What's this?

This is one of a set of introductory Go exercises by [John Arundel](https://bitfieldconsulting.com/golang/learn) called [For the Love of Go - Fundamentals](../README.md).


<small>Gopher image by [egonelbre](https://github.com/egonelbre/gophers)</small>