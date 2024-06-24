package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < len(numbers); i++ {
		fmt.Scan(&numbers[i])
	}
	last := len(numbers) - 1
	tmp := numbers[last]
	for i := last; i > 0; i-- {
		numbers[i] = numbers[i-1]
	}
	numbers[0] = tmp
	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i], " ")
	}
}
