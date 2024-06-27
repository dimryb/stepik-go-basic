package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	mas := make([][]int, n)
	for i := 0; i < n; i++ {
		mas[i] = make([]int, m)
		for j := 0; j < m; j++ {
			mas[i][j] = (i + 1) * (j + 1)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(mas[i][j], " ")
		}
		fmt.Println()
	}
}
