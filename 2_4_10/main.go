package main

import (
	"fmt"
)

func main() {
	var minutes int
	fmt.Scan(&minutes)
	hours := minutes / 60
	pureMinutes := minutes % 60
	fmt.Println(minutes, "мин - это", hours, "час", pureMinutes, "минут.")
}
