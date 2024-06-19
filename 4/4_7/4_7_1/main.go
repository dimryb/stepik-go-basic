package main

import "fmt"

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)

	for j := a; j <= b; j++ {
		cnt := 0
		for i := 1; i <= j; i++ {
			if j%i == 0 {
				cnt++
			}
		}
		if cnt >= k {
			fmt.Print(j, " ")
		}
	}
}
