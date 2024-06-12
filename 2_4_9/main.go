package main

import (
	"fmt"
)

func main() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)
	price := a*100 + b
	priceN := price * n
	rub := priceN / 100
	kop := priceN % 100
	fmt.Println(rub, kop)
}
