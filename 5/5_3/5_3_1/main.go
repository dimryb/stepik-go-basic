package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}
	double := false
outerLoop:
	for i := 0; i < n; i++ {
		for k := i + 1; k < n; k++ {
			if numbers[i] == numbers[k] && i != k {
				double = true
				break outerLoop
			}
		}
	}
	if double == true {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
