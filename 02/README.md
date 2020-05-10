# 02 - Getting in Shape

![](../img/hiking.svg)

The next thing to do is ensure that our Go code is formatted correctly. All Go code is formatted in a standard way, using a tool called `gofmt` (pronounced, in case you were wondering, 'go-fumpt'). In this exercise we'll see how to use it to check and reformat our code (if necessary).

The file `hello.go` in this directory contains a deliberate formatting mistake.

**TASK:** Run `gofmt` to check the formatting and find out what is wrong:

**`gofmt -d hello.go`**

```
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
```

`gofmt` is telling us what it wants to do. It wants to remove (`-`) the line that has the closing brace indented, and replace it (`+`) with a line where the brace is not indented.

We could make this change using our editor, but there's no need: `gofmt` can do it for us.

**TASK:** Run:

**`gofmt -w hello.go`**

This will reformat the file in place. You can check that it's worked by running `gofmt -d` on the file again. Now there should be no output, because it's already formatted correctly.

## Auto-format FTW

If you're using an editor that has Go support, such as Vim, Visual Studio Code, or GoLand, you can set up the editor to automatically run your code through `gofmt` every time you save it. This will help you avoid forgetting to format your code. Check the documentation for your editor to see how to set this up.

## Done?

When `gofmt -d hello.go` shows no output, you're done! Go on to the next exercise:

* [Fixing a Hole](../03/README.md)

## What's this?

This is one of a set of introductory Go exercises by [John Arundel](https://bitfieldconsulting.com/golang/learn) called [For the Love of Go - Fundamentals](../README.md).

<small>Gopher image by [egonelbre](https://github.com/egonelbre/gophers)</small>