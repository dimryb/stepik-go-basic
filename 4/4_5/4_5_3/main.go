package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	n := N
	for n > 0 {
		lastDight := n % 10
		fmt.Print(lastDight)
		n /= 10
	}
}
