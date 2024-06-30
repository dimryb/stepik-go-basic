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
	for _, char := range str {
		if char >= '0' && char <= '9' {
			fmt.Print(string(char), " ")
		}
	}
}
