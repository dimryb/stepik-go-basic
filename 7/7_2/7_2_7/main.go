package main

import (
	"fmt"
	"strconv"
)

func sumOfDigits(num int) int {
	strNum := strconv.Itoa(num)
	sum := 0
	for _, digit := range strNum {
		digitValue := int(digit - '0')
		sum += digitValue
	}
	return sum
}

func compareDigitSums(num1, num2 int) int {
	sum1 := sumOfDigits(num1)
	sum2 := sumOfDigits(num2)

	if sum1 > sum2 {
		return 1
	} else if sum2 > sum1 {
		return 2
	} else {
		return 0
	}
}

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)

	result := compareDigitSums(num1, num2)

	fmt.Println(result)
}
