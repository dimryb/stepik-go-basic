package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)
	reversedNum1 := reverseNumber(num1)
	reversedNum2 := reverseNumber(num2)
	sum := reversedNum1 + reversedNum2
	fmt.Println(sum)
}

func reverseNumber(n int) int {
	strNum := strconv.Itoa(n)
	runes := []rune(strNum)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reversedNum, _ := strconv.Atoi(string(runes))
	return reversedNum
}
