package main

import (
	"fmt"

	"example.com/go-course/phase01/ep12-packages-modules/solutions/02/sensor"
)

func main() {
	fmt.Println("sensor-01:", sensor.Status(30, 35))
	fmt.Println("sensor-02:", sensor.Status(28, 35))
}
