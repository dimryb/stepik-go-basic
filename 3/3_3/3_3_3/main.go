package main

import (
	"fmt"
)

func main() {
	var x1, x2, x3 int
	fmt.Scan(&x1, &x2, &x3)
	if x1 == x2 {
		if x2 == x3 {
			fmt.Println(3)
		} else {
			fmt.Println(2)
		}
	} else {
		if x1 == x3 {
			fmt.Println(2)
		} else {
			if x2 == x3 {
				fmt.Println(2)
			} else {
				fmt.Println(0)
			}
		}
	}
}
