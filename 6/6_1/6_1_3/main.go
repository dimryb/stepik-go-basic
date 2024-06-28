package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	text := scanner.Text()
	_ = scanner.Scan()
	kStr := scanner.Text()
	k, _ := strconv.Atoi(kStr)
	var newTextByte []byte
	textLen := len(text)
	for i := 0; i < textLen; i++ {
		if i != k-1 {
			newTextByte = append(newTextByte, text[i])
		}
	}
	finalString := string(newTextByte)
	fmt.Println(finalString)
}
