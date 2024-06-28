package main

import (
	"fmt"
	"strings"
)

func zero[T any]() T {
	var t T
	return t
}

/**
 * copied Double Linked List from double_linked_list.go
 */
type Cell[T any] struct {
	data T
	next *Cell[T]
	prev *Cell[T]
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

//////////////////////END DOUBLE LINKED LIST //////////////////////////

type Node[T any] struct {
	data        T
	left, right *Node[T]
}

func (node *Node[T]) displayIndented(indent string, depth int) string {
	result := strings.Repeat(indent, depth) + fmt.Sprintf("%v", node.data) + "\n"
	if node.left != nil {
		result += node.left.displayIndented(indent, depth+1)
	}
	if node.right != nil {
		result += node.right.displayIndented(indent, depth+1)
	}
	return result
}

func (node *Node[T]) preOrder() string {
	result := fmt.Sprintf("%v", node.data)
	if node.left != nil {
		result += " " + node.left.preOrder()
	}
	if node.right != nil {
		result += " " + node.right.preOrder()
	}
	return result
}

func (node *Node[T]) inOrder() string {
	result := ""
	if node.left != nil {
		result += node.left.inOrder()
	}
	result += fmt.Sprintf("%v ", node.data)
	if node.right != nil {
		result += node.right.inOrder()
	}
	return result
}

func (node *Node[T]) postOrder() string {
	result := ""
	if node.left != nil {
		result += node.left.postOrder()
	}
	if node.right != nil {
		result += node.right.postOrder()
	}
	result += fmt.Sprintf("%v ", node.data)
	return result
}

func (node *Node[T]) breadthFirst() string {
	queue := makeDoublyLinkedList[*Node[T]]()
	// Add the root node to the queue.
	queue.pushTop(node)

	result := ""
	// As long as the queue is not empty, do the following.
	for !queue.isEmpty() {
		// Dequeue the next node pointer.
		n := queue.dequeue()
		// Add that node’s data to the result string.
		result += fmt.Sprintf("%v", n.data)
		// If the node has children, add them to the queue.
		if n.left != nil {
			queue.enqueue(n.left)
		}
		if n.right != nil {
			queue.enqueue(n.right)
		}
		// If the queue is not empty, add a space to the result.
		if !queue.isEmpty() {
			result += " "
		}
	}

	return result
}

func buildTree() *Node[string] {
	aNode := Node[string]{"A", nil, nil}
	bNode := Node[string]{"B", nil, nil}
	cNode := Node[string]{"C", nil, nil}
	aNode.left = &bNode
	aNode.right = &cNode

	dNode := Node[string]{"D", nil, nil}
	eNode := Node[string]{"E", nil, nil}
	bNode.left = &dNode
	bNode.right = &eNode

	gNode := Node[string]{"G", nil, nil}
	eNode.left = &gNode

	fNode := Node[string]{"F", nil, nil}
	cNode.right = &fNode

	hNode := Node[string]{"H", nil, nil}
	fNode.left = &hNode

	iNode := Node[string]{"I", nil, nil}
	jNode := Node[string]{"J", nil, nil}
	hNode.left = &iNode
	hNode.right = &jNode

	return &aNode
}

func main() {
	rootNode := buildTree()
	fmt.Println("Indented Display: \n", rootNode.displayIndented("  ", 0))
	fmt.Println("Preorder:     ", rootNode.preOrder())
	fmt.Println("Inorder:      ", rootNode.inOrder())
	fmt.Println("Postorder:    ", rootNode.postOrder())
	fmt.Println("Breadth first:", rootNode.breadthFirst())
}
