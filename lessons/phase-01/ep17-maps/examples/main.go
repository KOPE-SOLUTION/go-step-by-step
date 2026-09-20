package main

import "fmt"

func main() {
	temperatures := map[string]float64{"sensor-01": 0}
	value, ok := temperatures["sensor-01"]
	fmt.Println(value, ok)
	value, ok = temperatures["missing"]
	fmt.Println(value, ok)
}
