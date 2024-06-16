package main

import (
	"fmt"
)

func main() {
	var month int
	fmt.Scan(&month)
	var days int
	switch month {
	case 2:
		days = 29
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	default:
		return
	}
	fmt.Println(days)
}
