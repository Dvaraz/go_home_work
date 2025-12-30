package hw2

// todo: Реализовать метод Next() для структуры Node !!!
// Реализовать метод подсчёта времени и расстояния  calcDistance( x Node, y Node) от узла X до узла Y для структуры Route.

import (
	"fmt"
)

type Node struct {
	name string
	time int
	dist float32
	next *Node
}

type Route struct {
	head   *Node
	length int
}

func (r *Route) insertAtHead(name string, time int, dist float32) {
	nodeHead := &Node{name, time, dist, nil}
	if r.head == nil {
		r.head = nodeHead
	} else {
		temp := r.head
		r.head = nodeHead
		nodeHead.next = temp
	}
	r.length += 1
}

func (n *Node) Next() *Node {
	return n.next
}

func (r *Route) calcDistance(x, y *Node) (int, float32, error) {
	if x == nil || y == nil {
		return 0, 0, fmt.Errorf("узлы не могут быть nil")
	}

	var posX, posY int
	var foundX, foundY bool

	current := r.head
	for i := 0; current != nil; i++ {
		if current == x {
			posX = i
			foundX = true
		}
		if current == y {
			posY = i
			foundY = true
		}
		current = current.next
	}

	if !foundX || !foundY {
		return 0, 0, fmt.Errorf("узлы не найдены в маршруте")
	}

	if posX > posY {
		return 0, 0, fmt.Errorf("узел x должен находиться раньше узла y в маршруте")
	}

	totalTime := 0
	totalDist := float32(0.0)

	current = x
	for current != y && current != nil {
		if current.next != nil {
			totalTime += current.next.time - current.time
			totalDist += current.next.dist - current.dist
		}
		current = current.next
	}

	if current != y {
		return 0, 0, fmt.Errorf("не удалось пройти от узла x до узла y")
	}

	return totalTime, totalDist, nil
}

func Task19() {
	list := Route{nil, 0}
	list.insertAtHead("С", 15, 18)
	list.insertAtHead("B", 5, 6)
	list.insertAtHead("A", 0, 0)
	fmt.Println(list.head)

	nodeA := list.head
	nodeB := nodeA.Next()
	nodeC := nodeB.Next()

	fmt.Printf("Следующий узел после %s: %s\n", nodeA.name, nodeA.Next().name)

	timeAC, distAC, err := list.calcDistance(nodeA, nodeC)
	if err == nil {
		fmt.Printf("От %s до %s: время=%d, расстояние=%.2f\n",
			nodeA.name, nodeC.name, timeAC, distAC)
	} else {
		fmt.Printf("Ошибка: %v\n", err)
	}
}
