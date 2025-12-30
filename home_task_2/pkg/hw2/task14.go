package hw2

import "fmt"

//todo: Дан массив размера N. Найти минимальное растояние между одинаковыми значениями в массиве и вывести их индексы.
// Одинаковых значение может быть два и более !
// Пример:
// var mass = [...]int{1, 2, 17, 54, 30, 89, 2, 1, 6, 2}

// Для числа 1 минимальное растояние в массиве по индексам: 0 и 7
// Для числа 2 минимальное растояние в массиве по индексам: 6 и 9
// Для числа 17 нет минимального растояния т.к элемент в массиве один.

func Task14(mass []int) {
	indexes := make(map[int][]int)

	for ind, value := range mass {
		indexes[value] = append(indexes[value], ind)
		if len(indexes[value]) > 2 {
			indexes[value] = indexes[value][1:]
		}
	}
	for key, value := range indexes {
		if len(value) > 1 {
			fmt.Printf("Для числа %d минимальное растояние в массиве по индексам: %d и %d\n", key, value[0], value[1])
		} else {
			fmt.Printf("Для числа %d нет минимального растояния т.к элемент в массиве один.\n", key)
		}
	}
}
