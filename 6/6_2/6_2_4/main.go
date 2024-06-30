package main

import "fmt"

func main() {
	var code string
	fmt.Scan(&code)
	char := rune(code[0])
	var newChar rune
	if char >= 'a' && char <= 'z' {
		newChar = char - 'a' + 'A'
	} else if char >= 'A' && char <= 'Z' {
		newChar = char - 'A' + 'a'
	} else {
		return
	}
	fmt.Println(string(newChar))
}
