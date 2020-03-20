A slight twist. This time, we have a `Greeting` function... but the test is missing!

```
--- FAIL: TestHello (0.00s)
    hello_test.go:8: not implemented yet
FAIL
exit status 1
FAIL    hello   0.073s
```

You'll need to modify `TestHello` so that it actually tests the `Greeting` function, in a similar way to the tests for previous exercises. Your test should call `hello.Greeting()`, assign the result to a variable, compare it to an expected result, and fail with `t.Errorf(...)` if they are different.

Note that the test package will need to import the `hello` package, in order to call `hello.Greeting`. Whenever we reference a Go package (like `testing`) we need to import it. Look at the `hello_test.go` file in previous exercises to see how to import `hello` so we can use it.

Once your test is passing, check that the test actually tests what you think it does. To do this, modify the `Greeting` function to return something that the test is not expecting.

The test should now fail. If it doesn't, it's not a very helpful test!

Fix the `Greeting` function so that it passes the test again, and go on to the next exercise.