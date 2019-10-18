package binarytree

import "fmt"

// Node struct represents a single node in the Tree
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// IsEmpty returns true if the Binary tree is empty
func (n *Node) IsEmpty() bool {
	return n == nil
}

// Insert an element in the correct spot in the tree
// If the tree is empty creates a new tree
func Insert(n *Node, data int) *Node {
	if n.IsEmpty() {
		return NewNode(data)
	} else if data < n.Data {
		n.Left = Insert(n.Left, data)
	} else if data > n.Data {
		n.Right = Insert(n.Right, data)
	}
	return n
}

// NewNode adds a new node in the empty slot
func NewNode(data int) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

// Build123a Implements 1-2-3 Binary Search tree in the first way
// by calling newNode() three times, and using three pointer variables
func Build123a() {
	// var root = new(Node)
	// root.Data = 2
	// root.Left = NewNode(1)
	// root.Right = NewNode(3)

	root := NewNode(2)
	leftNode := NewNode(1)
	rightNode := NewNode(3)

	root.Left = leftNode
	root.Right = rightNode
	root.PrintTree()
}

// Build123b Implements 1-2-3 Binary Search tree in the second way
// by calling newNode() three times, and using only one pointer variable
func Build123b() {
	var root = new(Node)
	root.Data = 2
	root.Left = NewNode(1)
	root.Right = NewNode(3)

	root.PrintTree()
}

// PrintTree prints the current tree
func (n *Node) PrintTree() {
	if n.IsEmpty() {
		return
	}
	n.Left.PrintTree()
	fmt.Print(n.Data, " ")
	n.Right.PrintTree()
}
