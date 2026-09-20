package main

import (
	"fmt"

	"example.com/go-course/basic/ep21/sensor"
)

func main() {
	fmt.Println(sensor.IsWarning(30))
	fmt.Println(sensor.IsWarning(31))
}
