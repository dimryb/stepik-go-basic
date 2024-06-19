package main

import (
	"fmt"
)

func main() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)
	nod := 1
	for i := 1; i <= n1 && i <= n2; i++ {
		if n1%i == 0 && n2%i == 0 {
			nod = i
		}
	}
	fmt.Println(nod)
}
