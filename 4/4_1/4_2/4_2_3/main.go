package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	sum := 0
	for i := 1; i <= x; i++ {
		var number int
		fmt.Scan(&number)
		if number%2 == 0 && number%3 != 0 {
			sum += number
		}
	}
	fmt.Println(sum)
}
