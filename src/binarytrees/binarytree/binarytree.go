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
func Build123a() *Node {
	// var root = new(Node)
	// root.Data = 2
	// root.Left = NewNode(1)
	// root.Right = NewNode(3)

	root := NewNode(2)
	leftNode := NewNode(1)
	rightNode := NewNode(3)

	root.Left = leftNode
	root.Right = rightNode

	return root
}

// Build123b Implements 1-2-3 Binary Search tree in the second way
// by calling newNode() three times, and using only one pointer variable
func Build123b() *Node {
	var root = new(Node)
	root.Data = 2
	root.Left = NewNode(1)
	root.Right = NewNode(3)

	return root
}

// Build123c Implements 1-2-3 Binary tree in the third way
// by calling Insert three times
func Build123c() *Node {
	var root = NewNode(2)
	root = Insert(root, 1)
	root = Insert(root, 3)

	return root
}

// Size calculates the total number of nodes in a binary tree
func (n *Node) Size() int {
	if n.IsEmpty() {
		return 0
	}
	return (n.Left.Size() + 1 + n.Right.Size())

}

// MaxDepth returns the maximum depth of the tree
func (n *Node) MaxDepth() int {
	if n.IsEmpty() {
		return 0
	}
	ldepth := n.Left.MaxDepth()
	rdepth := n.Right.MaxDepth()

	if ldepth < rdepth {
		return rdepth + 1
	}

	return ldepth + 1
}

// MinValue returns the smallest value in the binary tree
func (n *Node) MinValue() int {
	if n.IsEmpty() {
		fmt.Println("Tree is empty")
		return 0
	}

	if n.Left.IsEmpty() {
		return n.Data
	}

	return n.Left.MinValue()
}

// PrintTree prints the current tree in
func (n *Node) PrintTree() {
	if n.IsEmpty() {
		return
	}
	n.Left.PrintTree()
	fmt.Print(n.Data, " ")
	n.Right.PrintTree()
}

// PrintPostorder prints the tree in post order
func (n *Node) PrintPostorder() {
	if n.IsEmpty() {
		return
	}
	n.Left.PrintPostorder()
	n.Right.PrintPostorder()
	fmt.Print(n.Data, " ")
}

// HasPathSum returns true if the given binary tree has a branch
// whose elements add up to the given sum
func (n *Node) HasPathSum(sum int) bool {
	if n.IsEmpty() {
		return sum == 0
	}
	sum -= n.Data
	return n.Left.HasPathSum(sum) || n.Right.HasPathSum(sum)
}

// HasNoChildren returns true if the current node has no children, false
// otherwise
func (n *Node) HasNoChildren() bool {
	return n.Left.IsEmpty() && n.Right.IsEmpty()
}

// PrintPaths will print all the paths from the root to the leaf nodes
func (n *Node) PrintPaths() {
	var pathsize = make([]int, 0, n.MaxDepth())

	printPathsRecur(n, pathsize)
}

func printPathsRecur(n *Node, arr []int) {
	if n.IsEmpty() {
		return
	}
	// append this node to the path array
	arr = append(arr, n.Data)

	if n.HasNoChildren() {
		printArray(arr)
		return
	}
	// otherwise try both subtrees
	printPathsRecur(n.Left, arr)
	printPathsRecur(n.Right, arr)

}

// Utility that prints out an array on a line.
func printArray(arr []int) {
	fmt.Printf("\t")
	for _, val := range arr {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

// Mirror returns a mirror image of the given tree
// with the children of every node swapped left-to-right
func (n *Node) Mirror() *Node {

	if n.IsEmpty() {
		return nil
	}

	newTree := NewNode(n.Data)
	newTree.Left = n.Left.Mirror()
	newTree.Right = n.Right.Mirror()

	newTree.Left, newTree.Right = newTree.Right, newTree.Left
	return newTree
}

// DoubleTree adds a clone of each node as its left child in each subtree
func (n *Node) DoubleTree() {
	if n.IsEmpty() {
		return
	}
	clone := NewNode(n.Data)
	n.Left.DoubleTree()
	n.Right.DoubleTree()

	clone.Left, n.Left = n.Left, clone
}
