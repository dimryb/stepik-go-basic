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
		if prevous < n {
			cnt++
		}
		prevous = n
		fmt.Scan(&n)
	}
	fmt.Println(cnt)
}
