package main

import "fmt"

func showDevice(name string) {
	fmt.Println("Device:", name)
}

func main() {
	showDevice("sensor-01")
	showDevice("sensor-02")
}
