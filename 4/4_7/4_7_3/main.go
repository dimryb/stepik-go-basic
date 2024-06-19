package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	ost := N
	for i := 2; i <= ost; i++ {
		for ost%i == 0 {
			ost /= i
			fmt.Print(i, " ")
		}
	}
}
