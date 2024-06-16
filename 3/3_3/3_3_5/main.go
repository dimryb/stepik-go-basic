package main

import (
	"fmt"
)

func main() {
	var x, y float64
	fmt.Scan(&x, &y)
	if x > 0 {
		if y > 0 {
			fmt.Println(1)
		} else if y < 0 {
			fmt.Println(4)
		}
	} else if x < 0 {
		if y > 0 {
			fmt.Println(2)
		} else if y < 0 {
			fmt.Println(3)
		}
	}
}
