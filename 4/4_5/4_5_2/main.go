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
		lastDight := n % 10
		sum += lastDight
		n /= 10
	}

	if N%sum == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
