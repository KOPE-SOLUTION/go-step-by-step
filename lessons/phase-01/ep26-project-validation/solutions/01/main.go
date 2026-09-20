package main

import "fmt"

func main() {
	reading := Reading{DeviceID: "", Celsius: 25}
	err := validate(reading)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("%s: %.1f C [%s]\n", reading.DeviceID, reading.Celsius, status(reading))
}
