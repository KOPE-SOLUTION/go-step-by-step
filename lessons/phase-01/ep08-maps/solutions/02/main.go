package main

import "fmt"

func show(latest map[string]float64, deviceID string) {
	value, ok := latest[deviceID]
	if !ok {
		fmt.Println(deviceID + ": not found")
		return
	}
	fmt.Printf("%s: %.1f C\n", deviceID, value)
}

func main() {
	latest := map[string]float64{
		"sensor-01": 27.5,
		"sensor-02": 0,
	}
	latest["sensor-01"] = 29
	latest["sensor-03"] = 31
	show(latest, "sensor-01")
	show(latest, "sensor-02")
	show(latest, "sensor-04")
	delete(latest, "sensor-03")
	show(latest, "sensor-03")
	fmt.Println("devices:", len(latest))
}
