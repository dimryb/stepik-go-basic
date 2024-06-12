package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)
	a := number * number
	b := a * a
	c := a * b
	fmt.Println(c)
}
