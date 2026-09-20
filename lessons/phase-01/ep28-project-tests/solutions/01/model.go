package main

import (
	"errors"
	"fmt"
)

type Reading struct {
	DeviceID string
	Celsius  float64
}

func status(reading Reading) string {
	if reading.Celsius >= 30 {
		return "WARNING"
	}
	return "OK"
}

func validate(reading Reading) error {
	if reading.DeviceID == "" {
		return errors.New("device ID is empty")
	}
	if reading.Celsius < 0 || reading.Celsius > 100 {
		return errors.New("temperature outside simulated range")
	}
	return nil
}

func buildReport(readings []Reading) []string {
	lines := []string{}
	for _, reading := range readings {
		err := validate(reading)
		if err != nil {
			lines = append(lines, fmt.Sprintf("%s: ERROR: %v", reading.DeviceID, err))
			continue
		}
		line := fmt.Sprintf("%s: %.1f C [%s]", reading.DeviceID, reading.Celsius, status(reading))
		lines = append(lines, line)
	}
	return lines
}
