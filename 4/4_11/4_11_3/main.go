package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	out := 0
	index := 0
	for n > 0 {
		figure := n % 10
		if figure != 5 && figure != 7 {
			out += figure * int(math.Pow(10, float64(index)))
			index++
		}
		n /= 10
	}
	fmt.Println(out)
}
