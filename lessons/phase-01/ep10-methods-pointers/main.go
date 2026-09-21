package main

import "fmt"

type Reading struct {
	DeviceID string
	Celsius  float64
}

func (r Reading) Status() string {
	if r.Celsius >= 30 {
		return "WARNING"
	}
	return "OK"
}

func (r *Reading) Adjust(offset float64) {
	r.Celsius += offset
}

func adjustCopy(r Reading, offset float64) {
	r.Celsius += offset
}

func main() {
	reading := Reading{DeviceID: "sensor-01", Celsius: 29}
	adjustCopy(reading, 2)
	fmt.Printf("after copy: %.1f C [%s]\n", reading.Celsius, reading.Status())
	pointer := &reading
	pointer.Adjust(2)
	fmt.Printf("after pointer: %.1f C [%s]\n", reading.Celsius, reading.Status())
}
