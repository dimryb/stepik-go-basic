package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for n >= 10 {
		sum := 0
		for tmp := n; tmp > 0; tmp /= 10 {
			sum += tmp % 10
		}
		n = sum
	}

	fmt.Println(n)
}
