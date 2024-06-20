package main

import "fmt"

func main() {
	var N int
	for {
		fmt.Scan(&N)
		if N < 10 {
			continue
		}
		if N > 100 {
			break
		}
		fmt.Println(N)
	}
}
