package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	n := N
	for n > 0 {
		lastBit := n % 2
		fmt.Print(lastBit)
		n /= 2
	}
}
