package main

import "fmt"

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

func main() {
	readings := []Reading{
		{DeviceID: "sensor-01", Celsius: 27.5},
		{DeviceID: "sensor-02", Celsius: 30},
		{DeviceID: "sensor-03", Celsius: 31.5},
	}
	for _, reading := range readings {
		fmt.Printf("%s: %.1f C [%s]\n",
			reading.DeviceID, reading.Celsius, status(reading))
	}
	original := readings[0]
	copyReading := original
	copyReading.Celsius = 99
	fmt.Printf("original: %.1f, copy: %.1f\n", original.Celsius, copyReading.Celsius)
}
