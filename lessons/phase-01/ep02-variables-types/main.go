package main

import "fmt"

func main() {
	const unit = "C"
	var deviceID string = "sensor-01"
	celsius := 27.5
	connected := true
	var retries int
	var label string
	var failed bool
	celsius = 28.25

	fmt.Printf("%s: %.1f %s\n", deviceID, celsius, unit)
	fmt.Printf("connected=%t retries=%d\n", connected, retries)
	fmt.Printf("label=%q failed=%t\n", label, failed)
}
