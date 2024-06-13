package main

import (
	"fmt"
)

func main() {
	var F float64
	fmt.Scan(&F)
	deg := (F - 32) * 5 / 9
	fmt.Println(deg)
}
