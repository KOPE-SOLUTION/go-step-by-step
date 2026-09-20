package main

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		reading   Reading
		wantError bool
	}{
		{name: "lower", reading: Reading{DeviceID: "sensor-01", Celsius: 0}, wantError: false},
		{name: "upper", reading: Reading{DeviceID: "sensor-01", Celsius: 100}, wantError: false},
		{name: "below", reading: Reading{DeviceID: "sensor-01", Celsius: -0.1}, wantError: true},
		{name: "above", reading: Reading{DeviceID: "sensor-01", Celsius: 100.1}, wantError: true},
		{name: "empty ID", reading: Reading{Celsius: 25}, wantError: true},
	}
	for _, test := range tests {
		err := validate(test.reading)
		gotError := err != nil
		if gotError != test.wantError {
			t.Errorf("%s: error = %v; wantError %v", test.name, err, test.wantError)
		}
	}
}

func TestStatus(t *testing.T) {
	tests := []struct {
		input float64
		want  string
	}{
		{input: 29.9, want: "OK"},
		{input: 30, want: "WARNING"},
		{input: 30.1, want: "WARNING"},
	}
	for _, test := range tests {
		got := status(Reading{DeviceID: "sensor-01", Celsius: test.input})
		if got != test.want {
			t.Errorf("status(%v) = %q; want %q", test.input, got, test.want)
		}
	}
}

func TestBuildReport(t *testing.T) {
	readings := []Reading{
		{DeviceID: "sensor-01", Celsius: 30},
		{DeviceID: "broken", Celsius: -1},
		{DeviceID: "sensor-02", Celsius: 25},
	}
	want := []string{
		"sensor-01: 30.0 C [WARNING]",
		"broken: ERROR: temperature outside simulated range",
		"sensor-02: 25.0 C [OK]",
	}
	got := buildReport(readings)
	if len(got) != len(want) {
		t.Fatalf("line count = %d; want %d", len(got), len(want))
	}
	for index, line := range got {
		if line != want[index] {
			t.Errorf("line %d = %q; want %q", index, line, want[index])
		}
	}
}

func TestBuildReportEmpty(t *testing.T) {
	got := buildReport([]Reading{})
	if len(got) != 0 {
		t.Errorf("empty input: got %d lines; want 0", len(got))
	}
}

func TestReportTwoInvalid(t *testing.T) {
	got := buildReport([]Reading{
		{DeviceID: "bad-01", Celsius: -1},
		{DeviceID: "bad-02", Celsius: 101},
		{DeviceID: "good", Celsius: 25},
	})
	want := []string{
		"bad-01: ERROR: temperature outside simulated range",
		"bad-02: ERROR: temperature outside simulated range",
		"good: 25.0 C [OK]",
	}
	if len(got) != len(want) {
		t.Fatalf("got %d lines; want %d", len(got), len(want))
	}
	for index, line := range got {
		if line != want[index] {
			t.Errorf("line %d = %q; want %q", index, line, want[index])
		}
	}
}
