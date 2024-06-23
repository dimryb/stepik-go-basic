package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n < 1 {
		return
	}

	var x float64
	fmt.Scan(&x)

	minimum := x
	count := 1

	for i := 1; i < n; i++ {
		fmt.Scan(&x)
		if x < minimum {
			minimum = x
			count = 1
		} else if x == minimum {
			count++
		}
	}
	fmt.Println(count)
}
