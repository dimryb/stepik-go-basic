package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	var sign int
	if a < 0 {
		sign = -1
	} else if a > 0 {
		sign = 1
	} else {
		sign = 0
	}
	fmt.Println(sign)
}
