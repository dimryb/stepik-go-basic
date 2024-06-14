package main

import (
	"fmt"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	S := 1.0 / 2 * a * b
	fmt.Println(S)
}
