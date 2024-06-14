package main

import (
	"fmt"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	average := (a + b) / 2
	fmt.Println(average)
}
