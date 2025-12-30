package hw1

import "fmt"

// todo: Данные две переменные:
// age := 36.6
// temperature := 25
// Нужно обменять значения переменных местами. В итого age
// должен равнятся 25 а temperature – 36.6

func Task3() {
	age := 36.6
	temperature := 25.0

	age, temperature = temperature, age

	fmt.Println("temperature now is ", temperature, "age now is ", age)
}
