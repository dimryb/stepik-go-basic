package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	maximum := x
	var second int

	fmt.Scan(&x)
	if x > maximum {
		second = maximum
		maximum = x
	} else {
		second = x
	}

	for {
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		if x > maximum {
			second = maximum
			maximum = x
		} else if x > second && x < maximum {
			second = x
		}
	}

	fmt.Println(second)
}
