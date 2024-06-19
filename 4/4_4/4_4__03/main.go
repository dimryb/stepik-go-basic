package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)

	i := 1.0

	// fmt.Println("1:")
	// i = 1.0
	// for i < math.Sqrt(n) {
	// 	fmt.Println(i * i)
	// 	i++
	// }

	// fmt.Println("2:")
	// i = 1.0
	// for i*i < n {
	// 	fmt.Println(i * i)
	// 	i++
	// }

	fmt.Println("3:")
	i = 1.0
	for i*i <= n {
		fmt.Println(i * i)
		i++
	}

	// fmt.Println("4:")
	// i = 1.0
	// for 2*i == math.Sqrt(n) {
	// 	fmt.Println(i * i)
	// 	i++
	// }

	// fmt.Println("5:")
	// i = 1.0
	// for i == math.Sqrt(n) {
	// 	fmt.Println(i * i)
	// 	i++
	// }

	// fmt.Println("6:")
	// i = 1.0
	// for i*2 < n {
	// 	fmt.Println(i * i)
	// 	i++
	// }

	// fmt.Println("7:")
	// i = 1.0
	// for i*2 <= n {
	// 	fmt.Println(i * i)
	// 	i++
	// }

	fmt.Println("8:")
	i = 1.0
	for i <= math.Sqrt(n) {
		fmt.Println(i * i)
		i++
	}
}
