package hw2

import "fmt"

//todo: Заданы массивы

//Даны читатели книг
// var readers_books = [...] string {'id3', 'id5', 'id9', 'id8', 'id2', 'id1' }

//Даны читатели газет
// var readers_magazines = [...] string { 'id8', 'id2', 'id1', 'id4', 'id6', 'id7', 'id10'}

// Найти пользователей кто читает и книги и газеты

func Task18(readersBooks []string, readersMag []string) {
	var maxLenSlice *[]string
	var minLenSlice *[]string
	var readBoth []string
	longestReadersMap := make(map[string]bool)

	if len(readersBooks) == len(readersMag) {
		maxLenSlice = &readersBooks
		minLenSlice = &readersMag
	} else if len(readersBooks) > len(readersMag) {
		maxLenSlice = &readersBooks
		minLenSlice = &readersMag
	} else {
		maxLenSlice = &readersMag
		minLenSlice = &readersBooks
	}

	for _, reader := range *maxLenSlice {
		longestReadersMap[reader] = true
	}

	for _, reader := range *minLenSlice {
		if longestReadersMap[reader] {
			readBoth = append(readBoth, reader)
		}
	}
	fmt.Println(readBoth)
}
