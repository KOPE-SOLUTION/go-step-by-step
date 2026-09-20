package main

import "fmt"

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func main() {
	fmt.Printf("%.1f F\n", celsiusToFahrenheit(0))
}
