package main

import "fmt"

type Node struct {
	data        string
	left, right *Node
}

func (node *Node) insertValue(newData string) {
	newNode := Node{newData, nil, nil}
	var tempNode *Node = node
	for {
		if newData <= tempNode.data {
			if tempNode.left == nil {
				tempNode.left = &newNode
				break
			} else {
				tempNode = tempNode.left
			}
		} else {
			if tempNode.right == nil {
				tempNode.right = &newNode
				break
			} else {
				tempNode = tempNode.right
			}
		}
	}
}

func (node *Node) findValue(arg string) *Node {
	if node.data == arg {
		return node
	} else if node.data > arg && node.left != nil {
		return node.left.findValue(arg)
	} else if node.data < arg && node.right != nil {
		return node.right.findValue(arg)
	} else {
		return nil
	}
}

func (node *Node) inOrder() string {
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

func main() {
	// Make a root node to act as sentinel.
	root := Node{"", nil, nil}

	// Add some values.
	root.insertValue("I")
	root.insertValue("G")
	root.insertValue("C")
	root.insertValue("E")
	root.insertValue("B")
	root.insertValue("K")
	root.insertValue("S")
	root.insertValue("Q")
	root.insertValue("M")

	// Add F.
	root.insertValue("F")

	// Display the values in sorted order.
	fmt.Printf("Sorted values: %s\n", root.right.inOrder())

	// Let the user search for values.
	for {
		// Get the target value.
		target := ""
		fmt.Printf("String: ")
		fmt.Scanln(&target)
		if len(target) == 0 {
			break
		}

		// Find the value's node.
		node := root.findValue(target)
		if node == nil {
			fmt.Printf("%s not found\n", target)
		} else {
			fmt.Printf("Found value %s\n", target)
		}
	}
}
