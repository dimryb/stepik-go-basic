package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	mas := make([]int, n)
	var numbers []int
	for i := 0; i < n; i++ {
		fmt.Scan(&mas[i])
	}
	for i := 0; i < n; i++ {
		sign := false
		for _, value := range numbers {
			if value == mas[i] {
				sign = true
				break
			}
		}
		if sign == false {
			numbers = append(numbers, mas[i])
		}
	}
	fmt.Println(len(numbers))
}
