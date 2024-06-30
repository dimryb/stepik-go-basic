package main

import "fmt"

func main() {
	var str string
	var k int
	fmt.Scan(&str, &k)
	if k <= len(str) {
		fmt.Println(string(rune(str[k-1])))
	} else {
		fmt.Println("NO")
	}
}
