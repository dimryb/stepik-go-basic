package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	x1 := x % 10
	x2 := x / 10 % 10
	x3 := x / 100 % 10

	if x3 < x2 && x2 < x1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
