package main

import (
	"fmt"
)

func main() {
	var w int
	fmt.Scan(&w)

	if w%2 == 0 {
		w2 := w / 2
		if w2%2 == 0 {
			fmt.Print("YES")
			return
		} else {
			if w2 > 2 {
				fmt.Print("YES")
				return
			}
		}

	}

	fmt.Print("NO")
}
