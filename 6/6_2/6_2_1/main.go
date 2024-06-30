package main

import "fmt"

func main() {
	var code string
	fmt.Scan(&code)
	char := rune(code[0])
	if char >= '0' && char <= '9' {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
