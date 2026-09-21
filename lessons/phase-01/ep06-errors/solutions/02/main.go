package main

import (
	"fmt"
	"strconv"
)

func parseCelsius(text string) (float64, error) {
	if text == "" {
		return 0, fmt.Errorf("temperature is required")
	}
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0, fmt.Errorf("temperature must be a number: %q", text)
	}
	if value < 0 || value > 100 {
		return 0, fmt.Errorf("temperature outside simulated range: %.1f", value)
	}
	return value, nil
}

func show(text string) {
	value, err := parseCelsius(text)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("accepted: %.1f C\n", value)
}

func main() {
	show("27.5")
	show("")
	show("101")
	show("30")
}
