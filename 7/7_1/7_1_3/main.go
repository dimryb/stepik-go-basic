package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	z := digits(a) * digits(b)
	fmt.Println(z)
}

func digits(x int) int {
	counter := 0
	for x > 0 {
		x /= 10
		counter++
	}
	return counter
}
