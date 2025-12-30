package hw1

import (
	"fmt"
	"regexp"
	"strconv"
)

// todo: Преобразуйте переменную age и foo в число
// var := "23"
// foo := "23abc"
//
// Преобразуйте переменную age в Boolean
// age := "123abc"
//
// Преобразуйте переменную flag в Boolean
// flag := 1
//
// Преобразуйте значение в Boolean
// str_one := "Privet"
// str_two := ""
//
// Преобразуйте значение 0 и 1 в Boolean
//
// Преобразуйте false в строку

func Task2() {
	var1 := "23"
	foo := "23abc"

	num1, err := strconv.Atoi(var1)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	re_int := regexp.MustCompile(`^\d+`)
	match_int := re_int.FindString(foo)

	num2, err := strconv.Atoi(match_int)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(num1)
	fmt.Printf("var1 now is %T\n", num1)

	fmt.Println(num2)
	fmt.Printf("foo now is %T\n", num2)

	age := "123abc"

	bool1 := len(age) > 0

	fmt.Println(bool1)
	fmt.Printf("foo now is %T\n", bool1)

	flag := 1

	bool2 := flag != 0

	fmt.Println(bool2)
	fmt.Printf("flag now is %T\n", bool2)

	str_one := "Privet"
	str_two := ""

	bool3 := len(str_one) > 0
	bool4 := len(str_two) > 0

	fmt.Printf("bool3 = %t, bool4 = %t\n", bool3, bool4)
	fmt.Printf("bool3 now is %T, bool4 now is %T\n", bool3, bool4)

	bool5 := 1 != 0

	bool6 := 0 == 1

	fmt.Println(bool5, bool6)

	bool7 := false

	string_from_bool := strconv.FormatBool(bool7)

	fmt.Println(string_from_bool)
	fmt.Printf("boll7 now is %T\n", string_from_bool)

}
