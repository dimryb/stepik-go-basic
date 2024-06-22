package main

import "fmt"

func main() {
	start := 10
	for i := start; i < 100; i++ {
		x1 := i % 10
		x2 := i / 10 % 10
		if i == x1*x2*2 {
			fmt.Print(i)
		}
	}
}
