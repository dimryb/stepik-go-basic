package main

import (
	"fmt"
)

func main() {
	var n float64
	fmt.Scan(&n)
	cnt := 0
	sum := 0.0
	for n != 0 {
		cnt++
		sum += n
		fmt.Scan(&n)
	}
	mean := sum / float64(cnt)
	fmt.Println(mean)
}
