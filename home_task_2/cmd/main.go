package main

import (
	"hometask2/pkg/hw2"
)

func main() {
	// fmt.Println((hw2.Task11(12)))
	// fmt.Println((hw2.Task12(4, 1)))
	// hw2.Task12(2, 100)
	hw2.Task13(1905)
	// hw2.Task13(1947)
	// hw2.Task13(1988)
	// hw2.Task13(2001)
	// var mass = [...]int{1, 2, 17, 54, 30, 89, 2, 1, 6, 2}
	// hw2.Task14(mass[:])
	// arr := [100]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(hw2.Task15(arr[:]))
	// var allUsers = [...]string{"id3", "id5", "id9", "id8", "id2", "id1", "id4", "id6", "id7", "id10"}
	// var offlineUsers = [...]string{"id3", "id9", "id2", "id4", "id6", "id7"}
	// hw2.Task17(allUsers[:], offlineUsers[:])

	var readers_books = [...]string{"id3", "id5", "id9", "id8", "id2", "id1"}
	var readers_magazines = [...]string{"id8", "id2", "id1", "id4", "id6", "id7", "id10"}

	hw2.Task18(readers_books[:], readers_magazines[:])

	hw2.Task19()
}
