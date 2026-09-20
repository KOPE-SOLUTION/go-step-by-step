package main

import "testing"

func TestIsWarningBoundaries(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  bool
	}{
		{name: "below", input: 34.9, want: false},
		{name: "at threshold", input: 35, want: true},
		{name: "above", input: 35.1, want: true},
	}
	for _, test := range tests {
		got := isWarning(test.input)
		if got != test.want {
			t.Errorf("%s: got %v; want %v", test.name, got, test.want)
		}
	}
}
