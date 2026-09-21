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
	celsius := 30.0
	fmt.Printf("%.1f C = %.1f F [%s]\n",
		celsius, toFahrenheit(celsius), status(celsius, 30))
	fmt.Println("at threshold 35:", status(celsius, 35))
}
