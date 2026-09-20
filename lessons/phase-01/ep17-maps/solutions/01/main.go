package main

import "fmt"

func main() {
	temperatures := map[string]float64{}
	temperatures["sensor-02"] = 28.5
	value, ok := temperatures["sensor-02"]
	fmt.Println(value, ok)
}
