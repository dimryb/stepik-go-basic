package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	newStr := ""
	for _, char := range str {
		if char == 'e' {
			newStr += "i"
		} else {
			newStr += string(char)
		}
	}
	fmt.Println(newStr)
}
