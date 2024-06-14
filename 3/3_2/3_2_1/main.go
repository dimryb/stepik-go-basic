package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	if a < 1000 && a >= 100 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
