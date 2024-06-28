package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	text := scanner.Text()
	symbol := '\x00'
	for _, char := range text {
		if char == 'x' || char == 'w' {
			symbol = char
			break
		}
	}
	if symbol != '\x00' {
		fmt.Println(string(symbol))
	} else {
		fmt.Println("-1")
	}
}
