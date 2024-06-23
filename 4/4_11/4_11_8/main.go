package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := n + 1; ; i++ {
		x1 := i % 10
		x2 := i / 10 % 10
		x3 := i / 100 % 10
		x4 := i / 1000 % 10

		if x1 != x2 && x1 != x3 && x1 != x4 && x2 != x3 && x2 != x4 && x3 != x4 {
			fmt.Println(i)
			break
		}
	}
}
