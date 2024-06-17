package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	switch {
	case x <= 3:
		fmt.Println("начинающий")
	case x <= 7:
		fmt.Println("младший разработчик")
	case x <= 15:
		fmt.Println("средний разработчик")
	default:
		fmt.Println("старший разработчик")
	}
}
