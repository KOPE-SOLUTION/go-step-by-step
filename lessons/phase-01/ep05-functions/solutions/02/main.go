package main

import "fmt"

func toFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func status(celsius, threshold float64) string {
	if celsius >= threshold {
		return "WARNING"
	}
	return "OK"
}

func main() {
	for celsius := 29.0; celsius <= 31; celsius++ {
		fmt.Printf("%.0f C: %s\n", celsius, status(celsius, 30))
	}
}
