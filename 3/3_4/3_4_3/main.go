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
	x4 := x / 1000 % 10

	if x1 == x4 && x2 == x3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
