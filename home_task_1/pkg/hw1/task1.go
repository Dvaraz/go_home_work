package hw1

import "fmt"

// todo: Определить в коде переменные:

// 6. Задайте указатели на все перечисленные типы.
// Вывести их типы (надо погуглить).

func Task1() {
	// 1. Целочисленного типа
	var intVar int = 10
	var intPtr *int = &intVar

	// 2. Вещественного типа
	var floatVar float64 = 0.232
	var floatPtr *float64 = &floatVar

	// 3. Логического типа
	var boolVar bool = true
	var boolPtr *bool = &boolVar

	// 4. Строкового типа
	var stringVar string = "Hello"
	var stringPtr *string = &stringVar

	// 5. Пустого типа
	var nilVar interface{} = nil
	var nilPtr = &nilVar

	fmt.Printf("Type of intVar: %T\n", intPtr)
	fmt.Printf("Type of floatVar: %T\n", floatPtr)
	fmt.Printf("Type of boolVar: %T\n", boolPtr)
	fmt.Printf("Type of stringVar: %T\n", stringPtr)
	fmt.Printf("Type of nilVar: %T\n", nilPtr)
	fmt.Printf("Type of nilVar: %T\n", nilVar)
}
