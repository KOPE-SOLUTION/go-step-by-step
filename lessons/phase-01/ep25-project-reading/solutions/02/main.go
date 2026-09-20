package main

import "fmt"

func main() {
	reading := Reading{DeviceID: "sensor-03", Celsius: 29.96}
	fmt.Printf("%s: %.1f C [%s]\n", reading.DeviceID, reading.Celsius, status(reading))
}
