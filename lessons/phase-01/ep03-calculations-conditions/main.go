package main

import "fmt"

func main() {
	celsius := 30.0
	fahrenheit := celsius*9.0/5.0 + 32
	status := "OK"

	if celsius < 0 || celsius > 100 {
		status = "INVALID"
	} else if celsius >= 30 {
		status = "WARNING"
	}

	fmt.Printf("%.1f C = %.1f F [%s]\n", celsius, fahrenheit, status)
}
