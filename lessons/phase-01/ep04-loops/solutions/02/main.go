package main

import "fmt"

func main() {
	total := 0.0
	count := 0
	for round := 1; round <= 5; round++ {
		if round == 3 {
			fmt.Println("round 3: stopped")
			break
		}
		celsius := 25.0 + float64(round)
		total += celsius
		count++
		fmt.Printf("round %d: %.1f C\n", round, celsius)
	}
	if count > 0 {
		fmt.Printf("average: %.2f C (%d readings)\n", total/float64(count), count)
	}
}
