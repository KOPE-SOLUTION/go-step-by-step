package main

import (
	"errors"
	"fmt"
)

func validateTemperature(value float64) (float64, error) {
	if value < 0 || value > 100 {
		return 0, errors.New("temperature outside simulated range")
	}
	return value, nil
}

func main() {
	value, err := validateTemperature(100)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("Temperature:", value)
}
