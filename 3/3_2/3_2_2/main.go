package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	if x > -1 && x < 17 {
		fmt.Println("Принадлежит")
	} else {
		fmt.Println("Не принадлежит")
	}
}
