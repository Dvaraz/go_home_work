package hw2

// Создайте базовую структуру для персонажа из полей имени, уровня, опыта, класса
// Создайте структуру для хранения уникальных атрибутов Воина, Мага, Лучника с включением базовой структуры.
// Определите общие методы работы с персонажами в интерфейсе
// Реализуйте методы нанесения повреждений, воcстанновления здоровья, супер удара.
// Организуйте бой двух персонажей с нанесением удара и вычислением оставшегося здоровья.

type BaseUnit struct {
	Name       string
	level      int
	exp        float64
	profession string
}

// func NewUnit(class string) *BaseUnit {
// 	///
// }

type Warior struct {
	BaseStats BaseUnit
	health    int
	strength  int
}

type Mage struct {
	BaseStats BaseUnit
}

func Task16() {
	///
}
