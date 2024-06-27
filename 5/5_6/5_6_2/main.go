package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	mas := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&mas[i])
	}
	cnt := 0
	for i := 0; i < n-1; i++ {
		if mas[i] < mas[i+1] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
