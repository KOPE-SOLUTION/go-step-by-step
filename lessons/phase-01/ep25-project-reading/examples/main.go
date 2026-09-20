package main

import "fmt"

func main() {
	reading := Reading{DeviceID: "sensor-01", Celsius: 27.5}
	fmt.Printf("%s: %.1f C [%s]\n", reading.DeviceID, reading.Celsius, status(reading))
}
