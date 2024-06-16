package main

import (
	"fmt"
)

func main() {
	var k, m int
	fmt.Scan(&k, &m)
	days := k / m
	if k%m == 0 {
		fmt.Println(days)
	} else {
		fmt.Println(days + 1)
	}
}
