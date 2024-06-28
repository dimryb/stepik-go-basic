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
	count := 0
	for _, char := range text {
		if char == ' ' {
			count++
		}
	}
	words := count + 1
	fmt.Println(words)
}
