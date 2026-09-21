package main

import "fmt"

func main() {
	initial := [3]float64{27.5, 30.0, 28.5}
	readings := []float64{}
	for _, value := range initial {
		readings = append(readings, value)
	}
	readings = append(readings, 32.0)

	total := 0.0
	warnings := 0
	for index, value := range readings {
		total += value
		if value >= 30 {
			warnings++
		}
		fmt.Printf("%d: %.1f C\n", index, value)
	}
	if len(readings) > 0 {
		fmt.Printf("average: %.1f C, warnings: %d\n",
			total/float64(len(readings)), warnings)
	} else {
		fmt.Println("no readings")
	}
}
