package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)
	a1 := number % 10
	a2 := number % 100 / 10
	a3 := number % 1000 / 100
	fmt.Println(a1 + a2 + a3)
}
