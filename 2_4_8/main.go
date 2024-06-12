package main

import (
	"fmt"
)

func main() {
	var number1, number2, number3, number4 int
	fmt.Scan(&number1, &number2, &number3, &number4)
	fmt.Println(3 * (number1 + number2 + number3 + number4))
}
