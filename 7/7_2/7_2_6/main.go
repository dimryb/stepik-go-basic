package main

import (
	"fmt"
	"strconv"
)

func compareDigitCount(num1, num2 int) int {
	strNum1 := strconv.Itoa(num1)
	strNum2 := strconv.Itoa(num2)

	if len(strNum1) > len(strNum2) {
		return 1
	} else if len(strNum2) > len(strNum1) {
		return 2
	} else {
		return 0
	}
}

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)

	result := compareDigitCount(num1, num2)

	fmt.Println(result)
}
