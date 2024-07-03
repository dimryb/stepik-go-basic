package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if happyTicket(a) && happyTicket(b) {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}
}

func sum3(x int) int {
	sum := 0
	x %= 1000
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return sum
}

func happyTicket(x int) bool {
	return sum3(x) == sum3(x/1000)
}
