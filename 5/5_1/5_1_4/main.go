package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < len(numbers); i++ {
		fmt.Scan(&numbers[i])
	}

	end := len(numbers) / 2 * 2
	for i := 0; i < end; i += 2 {
		tmp := numbers[i]
		numbers[i] = numbers[i+1]
		numbers[i+1] = tmp
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i], " ")
	}
}
