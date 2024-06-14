package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var max int
	if a >= b {
		max = a
	} else {
		max = b
	}
	fmt.Println(max)
}
