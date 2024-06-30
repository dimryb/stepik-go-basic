package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str1 := scanner.Text()
	scanner.Scan()
	str2 := scanner.Text()
	str1Start := rune(str1[0])
	str2End := rune(str2[len(str2)-1])
	if str1Start == str2End {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
