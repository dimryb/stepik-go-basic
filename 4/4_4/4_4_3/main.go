package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	pokazatel := 0
	stepen := 1

	for stepen < n {
		stepen = stepen * 2
		pokazatel++
	}
	//fmt.Println("2 ^", pokazatel, "=", stepen)
	fmt.Println(pokazatel)
}
