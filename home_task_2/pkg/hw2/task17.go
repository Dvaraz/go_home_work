package hw2

import "fmt"

//todo: Заданы массивы
// все пользователи
// var all_users = [...]string {'id3', 'id5', 'id9', 'id8', 'id2', 'id1', 'id4', 'id6', 'id7', 'id10'}

// пользователи не в сети
// var offline_users = [...]string  {'id3', 'id9', 'id7', 'id2', 'id4', 'id6'}

// Вычислить пользователей online

func Task17(allUsers []string, offUsers []string) {
	var onlineUsers []string
	offlineUsersMap := make(map[string]bool)

	for _, user := range offUsers {
		offlineUsersMap[user] = true
	}

	for _, user := range allUsers {
		if !offlineUsersMap[user] {
			onlineUsers = append(onlineUsers, user)
		}
	}

	fmt.Println("Пользователи online: ", onlineUsers)
}
