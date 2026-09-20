package main

import "fmt"

func main() {
	reading := Reading{DeviceID: "sensor-01", Celsius: 100}
	err := validate(reading)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("%s: %.1f C [%s]\n", reading.DeviceID, reading.Celsius, status(reading))
}
