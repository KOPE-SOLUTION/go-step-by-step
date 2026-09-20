package main

import "fmt"

func isWarning(temperature float64) bool {
	return temperature >= 30
}

func main() {
	fmt.Println(isWarning(30))
}
