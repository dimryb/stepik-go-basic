package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	pokazatel := 0
	stepen := 1
	fmt.Println(stepen)
	for stepen*2 <= n {
		stepen = stepen * 2
		pokazatel++
		fmt.Println(stepen)
	}
}
