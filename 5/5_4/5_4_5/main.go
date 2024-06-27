package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	mas := make([][]int, n)
	sums := make([]int, n)
	for i := 0; i < n; i++ {
		mas[i] = make([]int, m)
		sums[i] = 0
		for j := 0; j < m; j++ {
			fmt.Scan(&mas[i][j])
			sums[i] += mas[i][j]
		}
	}
	maxInd := 0
	max := sums[0]
	for i := 1; i < n; i++ {
		if sums[i] > max {
			max = sums[i]
			maxInd = i
		}
	}
	fmt.Println(max)
	fmt.Println(maxInd)
}
