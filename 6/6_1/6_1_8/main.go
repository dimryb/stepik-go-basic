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
	newText := ""

	prevChar := rune(text[0])
	for _, char := range text {

		if !(char == ' ' && prevChar == ' ') {
			newText += string(char)
		}
		prevChar = char
	}

	fmt.Println(newText)
}
