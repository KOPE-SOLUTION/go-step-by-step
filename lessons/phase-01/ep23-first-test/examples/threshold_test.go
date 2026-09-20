package main

import "testing"

func TestIsWarning(t *testing.T) {
	got := isWarning(30)
	want := true
	if got != want {
		t.Fatalf("isWarning(30) = %v; want %v", got, want)
	}
}
