package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 0
	prevous := n
	for n != 0 {
		if prevous > 0 && n < 0 || prevous < 0 && n > 0 {
			cnt++
		}
		prevous = n
		fmt.Scan(&n)
	}
	fmt.Println(cnt)
}
