package main

import "fmt"

func main() {
	original := [2]int{1, 2}
	copyOfValues := original
	copyOfValues[0] = 9
	fmt.Println(original)
	fmt.Println(copyOfValues)
}
