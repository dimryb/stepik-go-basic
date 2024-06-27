package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var mas []int
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		if num%2 == 0 {
			mas = append(mas, num)
		}
	}
	for i := 0; i < len(mas); i++ {
		fmt.Print(mas[i], " ")
	}
	fmt.Println()
	fmt.Println(len(mas))
}
