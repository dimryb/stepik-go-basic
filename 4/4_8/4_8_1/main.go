package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	for i := N - 1; i > 1; i-- {
		if N%i == 0 {
			fmt.Println(i)
			break
		}
	}
}
