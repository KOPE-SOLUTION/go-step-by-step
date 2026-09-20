package main

import "fmt"

func addOffset(value *float64, offset float64) {
	*value = *value + offset
}

func main() {
	temperature := 20.0
	addOffset(&temperature, -0.5)
	fmt.Println(temperature)
}
