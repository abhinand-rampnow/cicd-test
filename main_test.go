package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(1, 3)
	want := 4

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
