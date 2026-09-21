package main

import "fmt"

func status(celsius, threshold float64) string {
	if celsius >= threshold {
		return "WARNING"
	}
	return "OK"
}

func main() {
	fmt.Println("29.9 C:", status(29.9, 30))
	fmt.Println("30.0 C:", status(30, 30))
}
