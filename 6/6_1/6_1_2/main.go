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
	var reverseTextByte []byte
	textLen := len(text)
	for i := 0; i < textLen; i++ {
		reverseTextByte = append(reverseTextByte, text[textLen-1-i])
	}
	finalString := string(reverseTextByte)
	fmt.Println(finalString)
}
