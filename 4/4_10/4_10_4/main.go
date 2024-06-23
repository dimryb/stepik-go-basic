package main

import "fmt"

func main() {
	var n int
	var x float64
	fmt.Scan(&n)
	if n < 1 {
		return
	}
	fmt.Scan(&x)
	maximum := x
	minimum := x
	for i := 1; i < n; i++ {
		fmt.Scan(&x)
		if x > maximum {
			maximum = x
		}
		if x < minimum {
			minimum = x
		}
	}
	var sign string
	if minimum < 30 {
		sign = "YES"
	} else {
		sign = "NO"
	}
	fmt.Println(maximum, sign)
}
