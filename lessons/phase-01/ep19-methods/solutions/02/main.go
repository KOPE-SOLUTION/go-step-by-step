package main

import "fmt"

type Reading struct {
	Celsius float64
}

func (r Reading) Fahrenheit() float64 {
	return r.Celsius*9/5 + 32
}

func main() {
	reading := Reading{Celsius: 25}
	fmt.Printf("%.1f F\n", reading.Fahrenheit())
}
