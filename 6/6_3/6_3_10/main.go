package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	sign := false
	for _, char := range s {
		if char == '-' {
			if sign {
				fmt.Print("2")
				sign = false
			} else {
				sign = true
			}
		} else if char == '.' {
			if sign {
				fmt.Print("1")
				sign = false
			} else {
				fmt.Print("0")
			}
		}
	}
}
