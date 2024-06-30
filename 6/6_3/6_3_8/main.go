package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	strings := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&strings[i])
	}
	for i := 0; i < n; i++ {
		str := strings[i]
		if len(str) > 10 {
			var strNumber string
			strNumber = fmt.Sprint(strNumber, len(str)-2)
			strings[i] = string(str[0]) + strNumber + string(str[len(str)-1])
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(strings[i])
	}
}
