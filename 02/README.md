The next thing to do is ensure that our Go code is formatted correctly. All Go code is formatted in a standard way, using a tool called `gofmt` (pronounced, in case you were wondering, 'go-fumpt').

The file `hello.go` in this directory contains a deliberate formatting mistake. We can use `gofmt` to check the formatting and tell us what is wrong:

```
gofmt -d hello.go
diff -u hello.go.orig hello.go
--- hello.go.orig       2020-03-19 20:43:37.000000000 +0000
+++ hello.go    2020-03-19 20:43:37.000000000 +0000
@@ -4,4 +4,4 @@
 // Greeting returns a friendly greeting message.
 func Greeting() string {
        return "Hello Gophers!"
-    }
+}
```

Look at the lines beginning with `-` and `+`:

```
-    }
+}
 }
```

`gofmt` is telling us what it wants to do. It wants to remove (`-`) the line that has the closing brace indented, and replace it (`+`) with a line where the brace is not indented.

To go ahead and make this change, run:

```
gofmt -w hello.go
```

If you're using an editor that has Go support, such as Vim, Visual Studio Code, or GoLand, you can set up the editor to automatically run your code through `gofmt` every time you save it. This will help you avoid forgetting to format your code. Check the documentation for your editor to see how to set this up.