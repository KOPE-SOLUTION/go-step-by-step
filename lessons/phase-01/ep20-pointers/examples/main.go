package main

import "fmt"

func addOffset(value *float64, offset float64) {
	*value = *value + offset
}

func main() {
	temperature := 25.0
	addOffset(&temperature, 0.5)
	fmt.Println(temperature)
}
