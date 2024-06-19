package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	sum := 0
	n := N
	for n > 0 {
		lastBit := n % 2
		if lastBit == 1 {
			sum++
		}
		n /= 2
	}
	fmt.Println(sum)
}
