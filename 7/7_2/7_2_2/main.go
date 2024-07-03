package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(average(a) + average(b))
}

func average(a int) float64 {
	sum := 0
	count := 0
	for i := 1; i <= a; i++ {
		sum += i
		count++
	}
	return float64(sum) / float64(count)
}
