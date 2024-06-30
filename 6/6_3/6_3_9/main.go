package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	countX := 0
	countDel := 0
	for i := 0; i < n; i++ {
		char := s[i]
		if char == 'x' {
			countX++
			if countX >= 3 {
				countDel++
			}
		} else {
			countX = 0
		}
	}
	fmt.Println(countDel)
}
