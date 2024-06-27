package main

import (
	"fmt"
	"strings"
)

type Cell[T any] struct {
	data T
	next *Cell[T]
	prev *Cell[T]
}

func zero[T any]() T {
	var t T
	return t
}

func (c *Cell[T]) addAfter(after *Cell[T]) {
	c.next = after
	after.prev = c
}

func (c *Cell[T]) addBefore(before *Cell[T]) {
	temp := c.prev

	temp.addAfter(before)
	before.addAfter(c)
}

func (c *Cell[T]) delete() {
	nextCell := c.next
	prevCell := c.prev

	if prevCell != nil {
		prevCell.next = nextCell
	}
	if nextCell != nil {
		nextCell.prev = prevCell
	}
	// set me to nil
	c = nil
}

type DoublyLinkedList[T any] struct {
	topSentinel    *Cell[T]
	bottomSentinel *Cell[T]
}

// Make a new DoublyLinkedList and initialize its topSentinel.
func makeDoublyLinkedList[T any]() DoublyLinkedList[T] {
	list := DoublyLinkedList[T]{}
	sentinel := &Cell[T]{zero[T](), nil, nil}
	list.topSentinel = sentinel
	list.bottomSentinel = sentinel

	list.bottomSentinel.next = list.topSentinel
	list.topSentinel.prev = list.bottomSentinel

	return list
}

func (l *DoublyLinkedList[T]) addRange(values []T) {
	var bottomCell *Cell[T] = l.bottomSentinel

	for _, v := range values {
		newCell := Cell[T]{data: v}
		bottomCell.addBefore(&newCell)
		bottomCell = newCell.prev.next
	}

}

func (l *DoublyLinkedList[T]) toString(separator string) string {
	buff := make([]string, 0)
	cell := l.topSentinel.next

	for cell != l.bottomSentinel {
		buff = append(buff, fmt.Sprintf("%v", cell.data))
		cell = cell.next
	}

	return strings.Join(buff, separator)
}

// Add an item to the top of the queue.
func (l *DoublyLinkedList[T]) enqueue(value T) {
	firstCell := l.topSentinel.next
	newCell := Cell[T]{data: value}
	firstCell.addBefore(&newCell)
}

// Remove an item from the bottom of the queue.
func (l *DoublyLinkedList[T]) dequeue() T {
	lastCell := l.bottomSentinel.prev
	if lastCell != l.topSentinel {
		res := lastCell.data
		lastCell.delete()
		return res
	}
	return zero[T]()
}

func (l *DoublyLinkedList[T]) pushBottom(value T) {
	newCell := Cell[T]{data: value}
	bottomSentinel := l.bottomSentinel
	bottomSentinel.addBefore(&newCell)
}

// Add an item at the top of the deque.
func (l *DoublyLinkedList[T]) pushTop(value T) {
	l.enqueue(value)
}

// Remove an item from the top of the deque.
func (l *DoublyLinkedList[T]) popTop() T {
	firstCell := l.topSentinel.next
	if firstCell != l.topSentinel {
		res := firstCell.data
		firstCell.delete()
		return res
	}
	return zero[T]()
}

// Add an item at the top of the deque.
func (l *DoublyLinkedList[T]) popBottom() T {
	return l.dequeue()
}

func (l *DoublyLinkedList[T]) length() int {
	len := 0
	cell := l.topSentinel.next

	for cell != nil {
		len++
		cell = cell.next
	}

	return len
}

func (l *DoublyLinkedList[T]) isEmpty() bool {
	return l.topSentinel.next == l.bottomSentinel
}

// func main() {
// 	// Make a list from a slice of values.
// 	list := makeDoublyLinkedList[string]()
// 	animals := []string{
// 		"Ant",
// 		"Bat",
// 		"Cat",
// 		"Dog",
// 		"Elk",
// 		"Fox",
// 	}
// 	list.addRange(animals)
// 	fmt.Println(list.toString(" "))
// }

func main() {
	// Test queue functions.
	fmt.Printf("*** Queue Functions ***\n")
	queue := makeDoublyLinkedList[string]()
	queue.enqueue("Agate")
	queue.enqueue("Beryl")
	fmt.Printf("%s ", queue.dequeue())
	queue.enqueue("Citrine")
	fmt.Printf("%s ", queue.dequeue())
	fmt.Printf("%s ", queue.dequeue())
	queue.enqueue("Diamond")
	queue.enqueue("Emerald")
	for !queue.isEmpty() {
		fmt.Printf("%s ", queue.dequeue())
	}
	fmt.Printf("\n\n")

	// Test deque functions. Names starting
	// with F have a fast pass.
	fmt.Printf("*** Deque Functions ***\n")
	deque := makeDoublyLinkedList[string]()
	deque.pushTop("Ann")
	deque.pushTop("Ben")
	fmt.Printf("%s ", deque.popBottom())
	deque.pushBottom("F-Cat")
	fmt.Printf("%s ", deque.popBottom())
	fmt.Printf("%s ", deque.popBottom())
	deque.pushBottom("F-Dan")
	deque.pushTop("Eva")
	for !deque.isEmpty() {
		fmt.Printf("%s ", deque.popBottom())
	}
	fmt.Printf("\n")
}
