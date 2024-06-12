package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scan(&input)
	number, _ := strconv.Atoi(input)
	fmt.Println(number * number)
}
