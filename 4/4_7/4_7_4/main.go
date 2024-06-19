package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	max := 0
	maxN := 0
	for i := 1; i <= N; i++ {
		k := i
		sum := 0
		for k != 0 {
			sum += k % 10
			k /= 10
		}
		if max < sum {
			max = sum
			maxN = i
		}
	}
	fmt.Println(maxN, max)
}
