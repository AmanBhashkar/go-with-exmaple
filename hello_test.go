package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("aman")
	want := "Hello aman"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
