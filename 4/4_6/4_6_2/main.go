package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	cntPositive := 0
	cntNegative := 0
	for n != 0 {
		if n > 0 {
			cntPositive++
		} else {
			cntNegative++
		}
		fmt.Scan(&n)
	}
	dif := cntPositive - cntNegative
	fmt.Println(dif)
}
