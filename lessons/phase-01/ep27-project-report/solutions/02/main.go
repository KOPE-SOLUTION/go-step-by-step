package main

import "fmt"

func main() {
	readings := []Reading{}
	for _, line := range buildReport(readings) {
		fmt.Println(line)
	}
}
