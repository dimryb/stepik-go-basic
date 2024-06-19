package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	sum := 0
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			sum++
		}
	}
	fmt.Println(sum)
}
