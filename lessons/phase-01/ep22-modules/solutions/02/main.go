package main

import (
	"fmt"

	"example.com/go-course/basic/ep22/sensor"
)

func main() {
	fmt.Println(sensor.IsWarning(35))
}
