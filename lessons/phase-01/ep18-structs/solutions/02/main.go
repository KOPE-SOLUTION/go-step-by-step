package main

import "fmt"

type Reading struct {
	DeviceID string
	Celsius  float64
}

func main() {
	original := Reading{DeviceID: "sensor-01", Celsius: 25}
	copied := original
	copied.Celsius = 30
	fmt.Println(original.Celsius, copied.Celsius)
}
