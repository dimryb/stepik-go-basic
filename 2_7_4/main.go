package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	hipotenuse := math.Sqrt(a*a + b*b)
	fmt.Println(hipotenuse)
}
