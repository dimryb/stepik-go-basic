package main

import "fmt"

func main() {
	var code string
	fmt.Scan(&code)
	char := rune(code[0])
	for i := 'a'; i <= char; i++ {
		fmt.Print(string(i), " ")
	}
}
