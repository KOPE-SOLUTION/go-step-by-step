package main

import "fmt"

func main() {
	readings := []Reading{{Celsius: 25}}
	for _, line := range buildReport(readings) {
		fmt.Println(line)
	}
}
