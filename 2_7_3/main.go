package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)
	char1 := number % 10
	char2 := number / 10 % 10
	char3 := number / 100 % 10
	sum := char1 + char2 + char3
	fmt.Println(sum)
}
