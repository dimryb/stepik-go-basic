package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	first := n % 10

	for n >= 10 {
		n /= 10
	}
	last := n

	if last == first {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
