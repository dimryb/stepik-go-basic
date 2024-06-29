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
	sign := true

	firstChar := rune(text[0])
	if !(firstChar >= 'a' && firstChar <= 'z' || firstChar >= 'A' && firstChar <= 'Z' || firstChar == '_') {
		sign = false
	}

	for i := 1; i < len(text); i++ {
		char := rune(text[i])
		if !(char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_' || char >= '0' && char <= '9') {
			sign = false
			break
		}
	}

	if sign {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
