package main

import "fmt"

func main() {
	temperature := 30.0
	if temperature >= 30 {
		fmt.Println("WARNING")
	} else {
		fmt.Println("OK")
	}
}
