package main

import "fmt"

func main() {
	var start, end string
	fmt.Scan(&start, &end)
	s, e := rune(start[0]), rune(end[0])
	if s > e {
		s, e = e, s
	}
	for i := s; i <= e; i++ {
		fmt.Print(string(i), " ")
	}
}
