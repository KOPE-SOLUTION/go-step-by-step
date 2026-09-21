package main

import (
	"fmt"

	"example.com/go-course/phase01/ep12-packages-modules/solutions/01/sensor"
)

func main() {
	fmt.Println("sensor-01:", sensor.Status(30, 30))
	fmt.Println("sensor-02:", sensor.Status(28, 30))
	fmt.Printf("30 C = %.1f F\n", sensor.ToFahrenheit(30))
}
