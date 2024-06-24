package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < len(numbers); i++ {
		fmt.Scan(&numbers[i])
	}

	for i := 0; i < len(numbers); i++ {
		number := numbers[i]
		if i%3 == 0 {
			fmt.Print(number, " ")
		}
	}
}
