package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	result := factorial(n) / (factorial(k) * factorial(n-k))
	fmt.Println(result)
}

func factorial(number int) int {
	fact := 1
	for i := 1; i <= number; i++ {
		fact = fact * i
	}
	return fact
}
