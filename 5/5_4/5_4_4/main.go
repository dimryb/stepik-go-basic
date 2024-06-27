package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	mas := make([][]int, n)
	for i := 0; i < n; i++ {
		mas[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				mas[i][j] = 1
			} else {
				mas[i][j] = mas[i-1][j] + mas[i][j-1]
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(mas[i][j], " ")
		}
		fmt.Println()
	}
}
