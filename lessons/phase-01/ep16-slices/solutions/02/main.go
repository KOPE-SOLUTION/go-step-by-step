package main

import "fmt"

func main() {
	values := []int{1, 2, 3}
	first := values[:2]
	first[0] = 9
	fmt.Println(values)
}
