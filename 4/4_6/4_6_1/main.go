package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	for n != 0 {
		if n%2 == 0 && n%3 != 0 {
			sum += n
		}
		fmt.Scan(&n)
	}
	fmt.Println(sum)
}
