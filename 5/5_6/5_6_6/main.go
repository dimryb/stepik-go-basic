package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	mas := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&mas[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mas[i]+mas[j] == 0 && mas[i] != 0 && mas[j] != 0 {
				fmt.Print(i, " ")
			}
		}
	}
}
