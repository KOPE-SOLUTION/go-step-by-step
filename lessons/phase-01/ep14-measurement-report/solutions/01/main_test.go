package main

import "testing"

func TestValidate(t *testing.T) {
	cases := []struct {
		name    string
		reading Reading
		wantErr bool
	}{
		{"normal", Reading{"sensor-01", 27.5}, false},
		{"empty ID", Reading{"", 27.5}, true},
		{"below range", Reading{"sensor-01", -0.1}, true},
		{"lower boundary", Reading{"sensor-01", 0}, false},
		{"upper boundary", Reading{"sensor-01", 100}, false},
		{"above range", Reading{"sensor-01", 100.1}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := validate(tc.reading)
			if (err != nil) != tc.wantErr {
				t.Errorf("validate(%+v) error = %v; wantErr %t", tc.reading, err, tc.wantErr)
			}
		})
	}
}

func TestBuildReportContinuesAfterInvalid(t *testing.T) {
	input := []Reading{
		{DeviceID: "sensor-01", Celsius: 29.9},
		{DeviceID: "sensor-02", Celsius: 101},
		{DeviceID: "sensor-03", Celsius: 30},
	}
	lines, accepted := buildReport(input, 30)
	want := []string{
		"sensor-01: 29.9 C [OK]",
		"sensor-02: ERROR: temperature outside simulated range",
		"sensor-03: 30.0 C [WARNING]",
	}
	if accepted != 2 || len(lines) != len(want) {
		t.Fatalf("accepted=%d lines=%v; want 2 accepted and %d lines", accepted, lines, len(want))
	}
	for i, line := range lines {
		if line != want[i] {
			t.Errorf("line %d = %q; want %q", i, line, want[i])
		}
	}
	if input[1].Celsius != 101 {
		t.Error("report changed the original measurement")
	}
}

func TestBuildReportEmpty(t *testing.T) {
	lines, accepted := buildReport(nil, 30)
	if len(lines) != 0 || accepted != 0 {
		t.Errorf("empty input: lines=%v accepted=%d; want no lines and zero accepted", lines, accepted)
	}
}

func TestBuildReportCustomThreshold(t *testing.T) {
	lines, accepted := buildReport([]Reading{{DeviceID: "sensor-01", Celsius: 30}}, 35)
	if accepted != 1 || len(lines) != 1 || lines[0] != "sensor-01: 30.0 C [OK]" {
		t.Errorf("custom threshold: lines=%v accepted=%d", lines, accepted)
	}
}
