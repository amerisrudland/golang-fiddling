package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ameris")
	want := "Hello, Ameris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
