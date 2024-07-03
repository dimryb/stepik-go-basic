package main

import (
	"fmt"
)

func maxElement(arr []int) int {
	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

func scanner(size int) []int {
	x := make([]int, size)
	for i := 0; i < size; i++ {
		var input int
		fmt.Scan(&input)
		x[i] = input
	}
	return x
}

func main() {
	var n, m int
	fmt.Scan(&n)
	a := scanner(n)

	fmt.Scan(&m)
	b := scanner(m)

	maxA := maxElement(a)
	maxB := maxElement(b)

	result := maxA * maxB
	fmt.Println(result)
}
