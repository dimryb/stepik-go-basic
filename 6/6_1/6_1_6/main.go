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
	for i := 0; i < len(text); i++ {
		for j := i + 1; j < len(text); j++ {
			if text[i] == text[j] {
				fmt.Println(string(text[i]))
				return
			}
		}
	}
}
