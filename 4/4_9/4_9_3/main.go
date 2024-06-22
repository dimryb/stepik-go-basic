package main

import "fmt"

func main() {
	cnt := 0
	for i := 100; i < 1000; i++ {
		x1 := i % 10
		x2 := i / 10 % 10
		x3 := i / 100 % 10
		if x1+x2+x3 == 8 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
