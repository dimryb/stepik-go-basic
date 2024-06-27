package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	mas := make([][]int, n)
	for i := 0; i < n; i++ {
		mas[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&mas[i][j])
		}
	}
	result := true
outerloop:
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if mas[i][j] != mas[j][i] {
				result = false
				break outerloop
			}
		}
	}
	if result {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
