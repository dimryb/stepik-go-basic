package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	maximum := a
	for i := a; i <= b; i++ {
		if i > maximum && i%7 == 0 {
			maximum = i
		}
	}
	if maximum%7 != 0 {
		fmt.Println("NO")
	} else {
		fmt.Println(maximum)
	}
}
