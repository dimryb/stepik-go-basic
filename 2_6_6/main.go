package main

import (
	"fmt"
	"math"
)

func main() {
	var bits float64
	fmt.Scan(&bits)
	bytes := bits / math.Pow(2, 13)
	fmt.Println(bytes)
}
