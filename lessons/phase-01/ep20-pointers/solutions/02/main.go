package main

import "fmt"

func main() {
	n := 5
	p := &n
	*p = *p + 1
	fmt.Println(n)
}
