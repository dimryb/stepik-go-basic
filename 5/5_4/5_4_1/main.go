package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	mas := make([][]int, n)
	for k := 0; k < n; k++ {
		mas[k] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var val int
			if i+j == n-1 {
				val = 1
			} else if i+j < n-1 {
				val = 0
			} else {
				val = 2
			}
			mas[i][j] = val
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(mas[i][j], " ")
		}
		fmt.Println()
	}
}
