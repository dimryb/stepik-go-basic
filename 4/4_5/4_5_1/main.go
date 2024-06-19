package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	for n > 0 {
		lastDight := n % 10
		if lastDight == 4 {
			sum++
		}
		n /= 10
	}
	fmt.Println(sum)
}
