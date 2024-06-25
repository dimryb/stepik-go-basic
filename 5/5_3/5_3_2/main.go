package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	var outNumbers []int
	for i := 0; i < n; i++ {
		double := false
		for k := 0; k < n; k++ {
			if numbers[i] == numbers[k] && i != k {
				double = true
				break
			}
		}
		if double == false {
			outNumbers = append(outNumbers, numbers[i])
		}
	}
	for _, number := range outNumbers {
		fmt.Print(number, " ")
	}
}
