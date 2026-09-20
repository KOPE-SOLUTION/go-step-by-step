package main

import "fmt"

func main() {
	const count = 2
	total := 5
	fmt.Println(total / count)
	average := float64(total) / count
	fmt.Printf("%.1f\n", average)
}
