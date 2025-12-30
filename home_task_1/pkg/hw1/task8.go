package hw1

import (
	"strconv"
)

// todo: Проверить истинность высказывания:
//"Данное четырехзначное число читается одинаково слева направо и справа налево".

func Task8(x int) bool {
	num := strconv.Itoa(x)
	return num == reverse(num)
}

func reverse(str string) string {
	var result string
	for _, k := range str {
		result = string(k) + result
	}
	return result
}
