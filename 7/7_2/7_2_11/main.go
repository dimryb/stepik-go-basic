package main

import (
	"fmt"
	"math"
)

// Функция для вычисления значения функции y(x)
func calculateY(x float64) float64 {
	if x < -5 {
		return 2*math.Abs(x) - 1
	} else if x >= -5 && x <= 5 {
		return x
	} else {
		return 2 * x
	}
}

func main() {
	sumY := 0.0

	// Вычисление суммы значений функции на интервале от -25 до 15 включительно с шагом 1
	for x := -25.0; x <= 15.0; x++ {
		sumY += calculateY(x)
	}

	fmt.Printf("Сумма значений функции на интервале от -25 до 15: %.2f\n", sumY)
}
