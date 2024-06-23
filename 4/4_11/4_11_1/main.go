package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 0
	for i := 0; i < n; i++ {
		var x float64
		fmt.Scan(&x)
		if x == 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
