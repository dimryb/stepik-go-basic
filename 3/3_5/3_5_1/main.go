package main

import (
	"fmt"
)

func main() {
	var age int
	var sex string
	fmt.Scan(&age, &sex)

	if sex == "m" && age >= 12 && age <= 18 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
