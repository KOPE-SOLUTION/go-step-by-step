package main

import "fmt"

func main() {
	temperature := 29.9
	if temperature >= 30 {
		fmt.Println("WARNING")
	} else {
		fmt.Println("OK")
	}
}
