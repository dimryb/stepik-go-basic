package main

import "fmt"

func main() {
	var n, h int
	fmt.Scan(&n, &h)
	mas := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&mas[i])
	}
	width := 0
	for i := 0; i < n; i++ {
		if mas[i] > h {
			width += 2
		} else {
			width += 1
		}
	}
	fmt.Println(width)
}
