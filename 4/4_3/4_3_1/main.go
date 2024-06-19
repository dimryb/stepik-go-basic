package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	mult := 1
	for i := 1; i <= N; i++ {
		mult *= i
	}
	fmt.Println(mult)
}
