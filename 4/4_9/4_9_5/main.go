package main

import "fmt"

func main() {
	cnt := 0
	for i := 1; i <= 1000; i++ {
		dividerCnt := 0
		for k := 1; k < i; k++ {
			if i%k == 0 {
				dividerCnt += k
			}
		}
		if i == dividerCnt {
			fmt.Print(i)
			cnt++
			if cnt < 3 {
				fmt.Print(",")
			} else {
				break
			}
		}
	}
}
