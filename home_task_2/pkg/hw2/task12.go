package hw2

import (
	"fmt"
	"strings"
)

// todo: Единицы массы пронумерованы следующим образом: 1 — килограмм, 2 — миллиграмм, 3 — грамм,
//  4 — тонна, 5 — центнер. Дан номер единицы массы и масса тела M в этих единицах (вещественное число).
//  Вывести массу данного тела в килограммах

func Task12(mesurment int, mass float64) string {
	massMesurment := map[int]float64{
		1: 1,
		2: 0.000001,
		3: 0.001,
		4: 1000,
		5: 50,
	}
	formated := fmt.Sprintf("%.6f", (massMesurment[mesurment] * mass))
	result := strings.TrimRight(formated, "0")
	result = strings.TrimRight(result, ".")
	return result
}
