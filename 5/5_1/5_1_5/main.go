package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < len(numbers); i++ {
		fmt.Scan(&numbers[i])
	}
	cnt := 0
	for i := 0; i < len(numbers)-1; i++ {
		x1 := numbers[i]
		x2 := numbers[i+1]
		if x1 > 0 && x2 > 0 || x1 < 0 && x2 < 0 {
			cnt++
		}
	}
	if cnt > 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
