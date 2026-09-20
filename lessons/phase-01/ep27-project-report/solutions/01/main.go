package main

import "fmt"

func main() {
	readings := []Reading{
		{DeviceID: "sensor-01", Celsius: 27.5},
		{DeviceID: "sensor-02", Celsius: 30},
		{DeviceID: "sensor-03", Celsius: -1},
		{DeviceID: "sensor-04", Celsius: 28},
		{DeviceID: "sensor-05", Celsius: 31},
	}
	for _, line := range buildReport(readings) {
		fmt.Println(line)
	}
}
