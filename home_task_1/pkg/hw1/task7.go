package hw1

import (
	"flag"
	"fmt"
)

//todo: Даны три точки A , B , C на числовой оси. Найти длины отрезков AC и BC и их сумму.
// Примечание: все точки получаем через консольный ввод.

func Task7() {
	A := flag.Int("a", 0, "точка A на числовой оси")
	B := flag.Int("b", 0, "точка B на числовой оси")
	C := flag.Int("c", 0, "точка C на числовой оси")

	flag.Parse()

	AC := Abs(*A - *C)
	BC := Abs(*B - *C)
	sum := AC + BC

	fmt.Println("Длинна отрезка AC=", AC, "Длинна отрезка BC=", BC)
	fmt.Println("Сумма отрезков AC и BC = ", sum)
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
