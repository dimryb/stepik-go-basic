package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}
	min := numbers[0]
	indexMin := 0
	max := numbers[0]
	indexMax := 0
	for index, number := range numbers {
		if min > number {
			min = number
			indexMin = index
		}
		if max <= number {
			max = number
			indexMax = index
		}
	}
	fmt.Println(indexMax - indexMin)
}
