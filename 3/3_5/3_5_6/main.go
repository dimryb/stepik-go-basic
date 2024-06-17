package main

import (
	"fmt"
)

func main() {
	var d1, d2, d3, d uint64
	fmt.Scan(&d1, &d2, &d3)
	if d1+d2 < d3 {
		d = d1*2 + d2*2
	} else if d1+d3 < d2 {
		d = d1*2 + d3*2
	} else if d3+d2 < d1 {
		d = d3*2 + d2*2
	} else {
		d = d1 + d2 + d3
	}
	fmt.Println(d)
}
