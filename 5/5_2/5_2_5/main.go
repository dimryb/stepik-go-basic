package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}
	polyndrom := true
	for i := 0; i < n/2; i++ {
		if numbers[i] != numbers[n-1-i] {
			polyndrom = false
			break
		}
	}
	if polyndrom {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
