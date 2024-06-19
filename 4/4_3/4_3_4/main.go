package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	mult := 1
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			mult *= i
		}
	}
	fmt.Println(mult)
}
