package main

import (
	"fmt"
)

func SumRange(X, Y int) int {
	if X > Y {
		return 0
	}
	sum := 0
	for i := X; i <= Y; i++ {
		sum += i
	}
	return sum
}

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	sumAB := SumRange(A, B)
	sumBC := SumRange(B, C)

	totalSum := sumAB + sumBC

	fmt.Println(totalSum)
}
