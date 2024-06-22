package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	cnt := 0
	for x := 0; x <= 1000; x++ {
		if x == e {
			continue
		}
		if a*x*x*x+b*x*x+c*x+d == 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
