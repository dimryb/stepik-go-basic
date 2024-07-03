package main

import (
	"fmt"
)

func countBs(s string) int {
	count := 0
	for _, char := range s {
		if char == 'b' { // Учитываем как 'b', так и 'B'
			count++
		}
	}
	return count
}

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)

	count1 := countBs(s1)
	count2 := countBs(s2)

	totalCount := count1 + count2
	fmt.Print(totalCount)
}
