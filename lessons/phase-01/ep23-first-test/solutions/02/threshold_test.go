package main

import "testing"

func TestIsWarningAtThreshold(t *testing.T) {
	got := isWarning(30)
	if got != true {
		t.Fatalf("isWarning(30) = %v; want true", got)
	}
}

func TestIsWarningAbove(t *testing.T) {
	got := isWarning(30.1)
	want := true
	if got != want {
		t.Fatalf("isWarning(30.1) = %v; want %v", got, want)
	}
}
