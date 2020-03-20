Now that you've written a function to pass a given test, and written a test for a given function, you can probably guess what's coming next. You're going to write both the test, and the function!

Create a file called `hello_test.go`, and write a test called TestHello, just like in previous exercises. Have it call the `Greeting()` function, and expect the result `"Howdy, Gopherinos!"`.

It should fail, because we haven't implemented `Greeting` yet:

```
go test
go build hello: no non-test Go files in exercises/06
FAIL    hello [build failed]
```

Create a `hello.go` file, just like the ones in previous exercises, and implement the `Greeting` function. Once your test passes, check the test again by modifying `Greeting` to return the wrong result, seeing the test fail, and then fixing `Greeting` so it passes again.
