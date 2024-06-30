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
	counter := 0
	for _, char := range str {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			counter++
		}
	}
	fmt.Println(counter)
}
