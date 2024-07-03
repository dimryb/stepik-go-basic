package main

import (
	"fmt"
)

func Fact2(N int) int {
	if N <= 0 {
		return 1
	}
	result := 1
	for i := N; i > 0; i -= 2 {
		result *= i
	}
	return result
}

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	factA := Fact2(A)
	factB := Fact2(B)
	factC := Fact2(C)

	fmt.Print(factA, " ")
	fmt.Print(factB, " ")
	fmt.Print(factC, " ")
}
