package hello_test

import (
	"hello"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello Gophers!"
	got := hello.Greeting()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
