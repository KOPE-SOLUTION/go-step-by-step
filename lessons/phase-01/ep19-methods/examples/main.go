package main

import "fmt"

type Reading struct {
	Celsius float64
}

func (r Reading) IsWarning() bool {
	return r.Celsius >= 30
}

func main() {
	reading := Reading{Celsius: 30}
	fmt.Println(reading.IsWarning())
}
