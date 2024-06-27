package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	mas := make([][]string, n)
	colors := []string{"C", "M", "Y"}
	for i := 0; i < n; i++ {
		mas[i] = make([]string, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&mas[i][j])
		}
	}
	colored := false
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for _, color := range colors {
				if mas[i][j] == color {
					colored = true
					break
				}
			}
		}
	}
	if colored {
		fmt.Println("#Color")
	} else {
		fmt.Println("#Black&White")
	}
}
