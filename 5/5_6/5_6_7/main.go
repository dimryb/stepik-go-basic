package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	mas := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&mas[i])
	}

	var taxis []int
	for i := 0; i < n; i++ {
		if mas[i] == 4 || mas[i] == 3 {
			taxis = append(taxis, mas[i])
			mas[i] = 0
		}
	}
	for i := 0; i < n; i++ {
		if mas[i] == 2 {
			sign := false
			for index, value := range taxis {
				if value == 2 {
					taxis[index] += mas[i]
					mas[i] = 0
					sign = true
					break
				}
			}
			if sign == false {
				taxis = append(taxis, mas[i])
				mas[i] = 0
			}
		}
	}
	for i := 0; i < n; i++ {
		if mas[i] == 1 {
			sign := false
			for index, value := range taxis {
				if value < 4 {
					taxis[index] += mas[i]
					mas[i] = 0
					sign = true
					break
				}
			}
			if sign == false {
				taxis = append(taxis, mas[i])
				mas[i] = 0
			}
		}
	}
	fmt.Println(len(taxis))
}
