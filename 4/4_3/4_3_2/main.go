package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	mult := 1
	for i := a; i <= b; i++ {
		mult *= i
	}
	fmt.Println(mult)
}
