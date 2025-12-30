package hw2

//todo: Дан целочисленный массив размера N из 10 элементов.
//Преобразовать массив, увеличить каждый его элемент на единицу.

func Task15(mass []int) []int {
	for i, v := range mass {
		if v > 0 {
			mass[i]++
		}
	}
	return mass
}
