package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Alejandro")
	want := "Hello, Alejandro"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
