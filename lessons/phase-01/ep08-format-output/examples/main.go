package main

import "fmt"

func main() {
	name := "sensor-01"
	temperature := 27.56
	online := true
	fmt.Printf("%s: %.1f C, online=%t\n", name, temperature, online)
}
