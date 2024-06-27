package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	mas := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&mas[i])
	}
	var k int
	fmt.Scan(&k)
	index := 0
	for i := 0; i <= n; i++ {
		if i == n {
			index = n + 1
			break
		}
		if mas[i] == k {
			index = i + 1
			continue
		}
		if mas[i] < k {
			index = i + 1
			break
		}
	}
	fmt.Println(index)
}
