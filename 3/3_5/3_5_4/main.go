package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)

	if x == 0 {
		fmt.Println("Обратного числа не существует")
	} else {
		_x := math.Pow(x, -1)
		fmt.Println(_x)
	}
}
