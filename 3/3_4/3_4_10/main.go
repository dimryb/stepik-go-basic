package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(a/2 + a%2 + b/2 + b%2 + c/2 + c%2)
}
