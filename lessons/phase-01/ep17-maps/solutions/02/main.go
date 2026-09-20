package main

import "fmt"

func main() {
	temperatures := map[string]float64{"sensor-01": 25}
	delete(temperatures, "sensor-01")
	value, ok := temperatures["sensor-01"]
	fmt.Println(value, ok)
}
