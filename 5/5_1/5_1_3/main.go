package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < len(numbers); i++ {
		fmt.Scan(&numbers[i])
	}

	count := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] < numbers[i] {
			count++
		}
	}
	fmt.Println(count)
}
