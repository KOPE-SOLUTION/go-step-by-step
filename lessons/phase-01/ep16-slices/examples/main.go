package main

import "fmt"

func main() {
	temperatures := []float64{27.5, 28}
	temperatures = append(temperatures, 30)
	fmt.Println(temperatures)
	fmt.Println("count:", len(temperatures))
}
