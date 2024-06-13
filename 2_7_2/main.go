package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)
	char3 := number / 100 % 10
	fmt.Println(char3)
}
