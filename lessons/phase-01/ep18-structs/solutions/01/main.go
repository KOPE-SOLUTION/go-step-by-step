package main

import "fmt"

type Reading struct {
	DeviceID string
	Celsius  float64
}

func main() {
	reading := Reading{DeviceID: "sensor-02"}
	fmt.Println(reading.DeviceID, reading.Celsius)
}
