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
	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-1-i] {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
