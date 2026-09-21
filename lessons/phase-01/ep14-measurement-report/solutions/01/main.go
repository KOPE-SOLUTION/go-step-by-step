package main

import "fmt"

type Reading struct {
	DeviceID string
	Celsius  float64
}

func validate(r Reading) error {
	if r.DeviceID == "" {
		return fmt.Errorf("device ID is required")
	}
	if r.Celsius < 0 || r.Celsius > 100 {
		return fmt.Errorf("temperature outside simulated range")
	}
	return nil
}

func status(celsius, threshold float64) string {
	if celsius >= threshold {
		return "WARNING"
	}
	return "OK"
}

func buildReport(readings []Reading, threshold float64) ([]string, int) {
	lines := []string{}
	accepted := 0
	for _, reading := range readings {
		err := validate(reading)
		if err != nil {
			lines = append(lines, fmt.Sprintf("%s: ERROR: %v", reading.DeviceID, err))
			continue
		}
		lines = append(lines, fmt.Sprintf("%s: %.1f C [%s]",
			reading.DeviceID, reading.Celsius, status(reading.Celsius, threshold)))
		accepted++
	}
	return lines, accepted
}

func main() {
	readings := []Reading{
		{DeviceID: "sensor-01", Celsius: 27.5},
		{DeviceID: "sensor-02", Celsius: 30},
		{DeviceID: "sensor-03", Celsius: 101},
		{DeviceID: "sensor-04", Celsius: 28},
		{DeviceID: "sensor-05", Celsius: 0},
	}
	lines, accepted := buildReport(readings, 30)
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Printf("accepted: %d/%d\n", accepted, len(readings))
}
