package main

import "fmt"

func main() {
	devices := []string{}
	devices = append(devices, "sensor-01")
	fmt.Println(devices)
	fmt.Println(len(devices))
}
