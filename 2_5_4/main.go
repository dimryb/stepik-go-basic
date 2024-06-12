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
	number2 := a1*100 + a2*10 + a3
	fmt.Println(number2)
}
