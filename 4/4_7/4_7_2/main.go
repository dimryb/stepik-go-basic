package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	cnt := 0
	for i := 1; i <= N; i++ {
		k := i
		for k != 0 {
			if k%10 == 7 {
				cnt++
			}
			k /= 10
		}
	}
	fmt.Println(cnt)
}
