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
	for i := 0; i < n; i++ {
		if mas[i]%3 == 0 && mas[i]%10 == 7 {
			cnt++
		}
	}
	for i := 0; i < n; i++ {
		if mas[i]%3 == 0 && mas[i]%10 == 7 {
			mas[i] = cnt
		}
	}
	for i := 0; i < n; i++ {
		fmt.Print(mas[i], " ")
	}
}
