package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	if x%2 == 1 {
		fmt.Println("YES")
	} else {
		if x >= 2 && x <= 5 {
			fmt.Println("NO")
		} else if x >= 6 && x <= 20 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
