package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	for i := 0; i <= 1000; i++ {
		if a*i*i*i+b*i*i+c*i+d == 0 {
			fmt.Print(i, " ")
		}
	}
}
