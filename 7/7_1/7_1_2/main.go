package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	z := sign(a) + sign(b)
	fmt.Println(z)
}

func sign(x int) int {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	default:
		return x
	}
}
