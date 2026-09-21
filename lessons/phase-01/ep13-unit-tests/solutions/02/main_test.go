package main

import "testing"

func TestStatus(t *testing.T) {
	cases := []struct {
		name      string
		celsius   float64
		threshold float64
		want      string
	}{
		{"below", 29.9, 30, "OK"},
		{"equal", 30, 30, "OK"},
		{"above", 30.1, 30, "WARNING"},
		{"custom threshold", 30, 35, "OK"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := status(tc.celsius, tc.threshold)
			if got != tc.want {
				t.Errorf("status(%v, %v) = %q; want %q",
					tc.celsius, tc.threshold, got, tc.want)
			}
		})
	}
}
