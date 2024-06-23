package main

import (
	"fmt"
)

func main() {
	var a, n int
	fmt.Scan(&a, &n)
	out := 1
	for i := 0; i < n; i++ {
		out *= a
	}
	fmt.Println(out)
}
