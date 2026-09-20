package main

import "fmt"

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func main() {
	fahrenheit := celsiusToFahrenheit(25)
	fmt.Printf("%.1f F\n", fahrenheit)
}
