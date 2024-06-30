package main

import "fmt"

func main() {
	var code string
	fmt.Scan(&code)
	char := rune(code[0])
	for i := char; i <= 'z'; i++ {
		fmt.Print(string(i), " ")
	}
}
