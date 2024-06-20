package main

import "fmt"

func main() {
	var s1 string
	fmt.Scan(&s1)
	cnt := 0
	for {
		var s2 string
		fmt.Scan(&s2)
		cnt++
		if s2 == s1 {
			fmt.Println(cnt)
			break
		}
	}
}
