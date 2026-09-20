package main

import "fmt"

func main() {
	temperatures := [3]float64{27.5, 28, 30}
	fmt.Println("count:", len(temperatures))
	for index, value := range temperatures {
		fmt.Println(index, value)
	}
}
