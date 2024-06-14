package main

import (
	"fmt"
)

func main() {
	var years int
	fmt.Scan(&years)
	if years < 14 {
		fmt.Println("детство")
	} else if years < 25 {
		fmt.Println("молодость")
	} else if years < 60 {
		fmt.Println("зрелость")
	} else {
		fmt.Println("старость")
	}
}
