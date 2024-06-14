package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	separator := scanner.Text()
	_ = scanner.Scan()
	line1 := scanner.Text()
	_ = scanner.Scan()
	line2 := scanner.Text()
	_ = scanner.Scan()
	line3 := scanner.Text()

	fmt.Print(line1, separator, line2, separator, line3)
}
