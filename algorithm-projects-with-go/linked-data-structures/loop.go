package main

import (
	"fmt"
	"strings"
)

type Cell[T any] struct {
	data T
	next *Cell[T]
}

func zero[T any]() T {
	var t T
	return t
}

func (c *Cell[T]) addAfter(after *Cell[T]) {
	c.next = after
}

func (c *Cell[T]) deleteAfter() *Cell[T] {
	if c.next == nil {
		panic("No pointer after to delete")
	}
	nextCell := c.next
	if nextCell.next != nil {
		c.next = nextCell.next
	} else {
		c.next = nil
	}
	return nextCell
}

type LinkedList[T any] struct {
	sentinel *Cell[T]
}

// Make a new LinkedList and initialize its sentinel.
func makeLinkedList[T any]() LinkedList[T] {
	list := LinkedList[T]{}
	list.sentinel = &Cell[T]{zero[T](), nil}
	return list
}

func (l *LinkedList[T]) addRange(values []T) {
	var lastCell *Cell[T] = l.sentinel
	for lastCell.next != nil {
		lastCell = lastCell.next
	}

	for _, v := range values {
		newCell := Cell[T]{data: v}
		lastCell.addAfter(&newCell)
		lastCell = lastCell.next
	}

}

func (l *LinkedList[T]) toString(separator string) string {
	buff := make([]string, 0)
	cell := l.sentinel.next

	for cell != nil {
		buff = append(buff, fmt.Sprintf("%v", cell.data))
		cell = cell.next
	}

	return strings.Join(buff, separator)
}

func (l *LinkedList[T]) toStringMax(separator string, max int) string {
	buff := make([]string, 0)
	cell := l.sentinel.next
	i := 0

	for cell != nil {
		buff = append(buff, fmt.Sprintf("%v", cell.data))
		cell = cell.next
		i++
		if i >= max {
			break
		}
	}

	return strings.Join(buff, separator)
}

func (l *LinkedList[T]) length() int {
	len := 0
	cell := l.sentinel.next

	for cell != nil {
		len++
		cell = cell.next
	}

	return len
}

func (l *LinkedList[T]) isEmpty() bool {
	return l.sentinel.next == nil
}

func (l *LinkedList[T]) push(t T) {
	var lastCell *Cell[T] = l.sentinel
	for lastCell.next != nil {
		lastCell = lastCell.next
	}
	newCell := Cell[T]{data: t}
	lastCell.addAfter(&newCell)
}

func (l *LinkedList[T]) pop() T {
	if l.isEmpty() {
		panic("Cannot pop an empty cell")
	}
	var penultimateCell, lastCell *Cell[T]

	penultimateCell = l.sentinel
	lastCell = penultimateCell.next
	if lastCell != nil {
		for lastCell.next != nil {
			penultimateCell = lastCell
			lastCell = penultimateCell.next
		}
	}
	var t T = lastCell.data
	penultimateCell.deleteAfter()
	return t
}

func (l *LinkedList[T]) hasLoop() bool {
	if l.isEmpty() {
		return false
	}
	length := l.length()
	slow := l.sentinel.next
	var fast *Cell[T]

	println(slow)
	println(fast)

	return slow == fast

	for {
		// 	if slow.next != nil && slow.next.next != nil {
		// 		slow = slow.next
		// 		fast = slow.next.next
		// 	} else {
		// 		return false
		// 	}
		// 	if slow == fast {
		// 		return true
		// 	}
	}
}

func main() {
	// Make a list from an array of values.
	values := []string{
		"0", "1", "2", "3", "4", "5",
	}
	list := makeLinkedList[string]()
	list.addRange(values)

	fmt.Println(list.toString(" "))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 5 point to cell 2.
	list.sentinel.next.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.toStringMax(" ", 10))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 4 point to cell 2.
	list.sentinel.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.toStringMax(" ", 10))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
}
