package main

import "fmt"

func main() {
	readings := []Reading{{DeviceID: "bad-01", Celsius: -1}, {DeviceID: "bad-02", Celsius: 101}, {DeviceID: "good", Celsius: 25}}
	for _, line := range buildReport(readings) {
		fmt.Println(line)
	}
}
