package main

import "fmt"

func main() {
	var f, i, o string
	fmt.Scan(&f, &i, &o)
	shortI := string(i[0]) + "."
	shortO := string(o[0]) + "."
	fmt.Print(f, " ", shortI, shortO)
}
