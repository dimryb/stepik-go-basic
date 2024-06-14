package main

import (
	"fmt"
)

func main() {
	var realNumber float64
	fmt.Scan(&realNumber)
	integer := int32(realNumber)
	fraction := realNumber - float64(integer)
	fmt.Println(fraction)
}
