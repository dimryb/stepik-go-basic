package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	sum := a + b + c
	comp := a * b * c
	fmt.Println(sum, comp)
}
