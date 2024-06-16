package main

import (
	"fmt"
)

func main() {
	var k2, k3, k5, k6 int
	fmt.Scan(&k2, &k3, &k5, &k6)

	var num32, num256 int

	if k2 <= k5 && k2 <= k6 {
		num256 = k2
	} else if k5 <= k2 && k5 <= k6 {
		num256 = k5
	} else {
		num256 = k6
	}

	k2 -= num256

	if k2 <= k3 {
		num32 = k2
	} else {
		num32 = k3
	}

	maxSum := num256*256 + num32*32
	fmt.Println(maxSum)
}
