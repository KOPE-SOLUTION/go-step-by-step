package main

import "testing"

func TestIsWarningAtThreshold(t *testing.T) {
	got := isWarning(30)
	if got != true {
		t.Fatalf("isWarning(30) = %v; want true", got)
	}
}

func TestIsWarningBelow(t *testing.T) {
	got := isWarning(29.9)
	want := false
	if got != want {
		t.Fatalf("isWarning(29.9) = %v; want %v", got, want)
	}
}
