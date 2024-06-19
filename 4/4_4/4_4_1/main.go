package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)

	sum := 0
	for a%3 == 0 {
		a /= 3
		sum++
	}

	fmt.Println(sum)
}
