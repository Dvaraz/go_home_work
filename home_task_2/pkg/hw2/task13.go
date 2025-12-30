package hw2

import (
	"fmt"
	"strings"
)

// В восточном календаре принят 60-летний цикл, состоящий из 12- летних подциклов,
// обозначаемых названиями цвета: зеленый, красный, желтый, белый и черный.
// В каждом подцикле годы носят названия животных: крысы, коровы, тигра, зайца, дракона,
// змеи, лошади, овцы, обезьяны, курицы, собаки и свиньи. По номеру года вывести его название,
// если 1984 год был началом цикла — годом зеленой крысы.

// https://clck.ru/3QWynz
func Task13(year int) {
	pivot := 1984
	animalMap := map[int]string{
		0:  "крысы",
		1:  "быка",
		2:  "тигра",
		3:  "кролика",
		4:  "дракона",
		5:  "змеи",
		6:  "лощади",
		7:  "козы",
		8:  "обезьяны",
		9:  "петуха",
		10: "собаки",
		11: "свиньи",
	}

	colorMap := map[int]string{
		0: "зелен",
		1: "красн",
		2: "желт",
		3: "бел",
		4: "черн",
	}
	year_animal := modulo(year-pivot, 12)
	year_color := modulo(year-pivot, 10) / 2
	animal := animalMap[year_animal]
	color := colorMap[year_color]
	ending := findMaleFemale(animal)

	fmt.Printf("%v год  это год %v%v %v\n", year, color, ending, animal)
}

func modulo(a, b int) int {
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}

func findMaleFemale(word string) string {
	male := "а"
	ending := "ой"
	if strings.HasSuffix(word, male) {
		ending = "ого"
	}
	return ending
}
