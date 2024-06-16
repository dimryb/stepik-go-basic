package main

import (
	"fmt"
)

func main() {
	var x, y int
	var symbol string
	var result float64

	fmt.Scan(&x, &y, &symbol)
	switch symbol {
	case "+":
		result = float64(x + y)
	case "-":
		result = float64(x - y)
	case "*":
		result = float64(x * y)
	case "/":
		if y != 0 {
			result = float64(x) / float64(y)
		} else {
			fmt.Println("На ноль делить нельзя!")
			return
		}
	default:
		fmt.Println("Неверная операция")
		return
	}

	fmt.Println(result)
}
