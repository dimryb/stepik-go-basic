package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	mult := 1
	for i := a; i <= b; i++ {
		if i%10 == 7 || i%10 == -7 {
			mult *= i
		}
	}
	fmt.Println(mult)
}
